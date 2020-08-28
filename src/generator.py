import sys, string, time
from datetime import datetime
from generation import core
from inputs.input_aggregator_names import aggregator_names
from inputs.input_aggregators import aggregators
from inputs.input_features import features
from inputs.input_patterns import patterns
# VS Code's Python 3, PyLint will report errors for a few of the imports above. They have not actually caused errors for me, they have been ignored. 


class CodeGeneratorBackend:

    def begin(self, tab="\t", spacer='-'):
        self.code = []
        self.tab = tab
        self.spacer = spacer
        self.level = 0

    def end(self):
        return "".join(self.code)

    def write(self, string):
        self.code.append(self.tab * self.level + string)

    def writeComment(self, string):
        self.writeLine('// ' + string)


    def writeLine(self, string):
        self.write(string)
        self.write('\n')

    def writeTitle(self, string):
        self.writeComment(len(string) * self.spacer)
        self.writeComment(string)
        self.writeComment(len(string) * self.spacer)
        c.newLine(1)

    def writeSubTitle(self, string):
        self.writeComment(3 * self.spacer)
        self.writeComment(string)
        self.writeComment(3 * self.spacer)

    def newLine(self, number):
        for i in list(range(0, number)):
            self.write('\n')

    def indent(self):
        self.level = self.level + 1

    def dedent(self):
        if self.level == 0:
            raise SyntaxError ("internal error in code generator")
        self.level = self.level - 1

    def writeHeader(self, patternName, featureName, aggregatorName):
        self.writeLine('func ' + aggregatorName + '_' + featureName + '_' + patternName + '(data []float64) float64 {')
        self.indent()

    def writeEntryState(self, patternName):
        res = 'currentState := \''
        res = res + str(core.getInitState(patternName))
        res = res + '\'\n'
        self.write(res)

    def writeInitValue(self, accumulator, patternName, featureName, aggregatorName):
        res = accumulator + ' := ' #res = accumulator + ' = '
        res = res + str(core.getInitValue(accumulator, patternName, featureName, aggregatorName))
        res = res + '\n'
        self.write(res)

    def writeFunction(self, patternName, featureName, aggregatorName, x):
        ## Stop production of extra 'Max/Min_one_' functions
        isMaxorMin = True if x == 'Max' or x == 'Min' else False
        isOne = True if featureName == 'one' else False
        if (isMaxorMin and isOne):
            return
        ## Stop production of 15 extra width functions
        isWidth = True if featureName == 'width' else False
        isDec = True if patternName == 'decreasing' else False
        isInc = True if patternName == 'increasing' else False
        isDip = True if patternName == 'dip_on_increasing_sequence' else False
        isBump = True if patternName == 'bump_on_decreasing_sequence' else False
        isOnlySteady = True if patternName == 'steady' else False
        if (isWidth and (isInc or isDec or isDip or isBump or isOnlySteady)):
            return
        ## Stop generation of 30 extra '_max_' functions
        isMaxFeature = True if featureName == 'max' else False
        isTerrace = True if patternName == 'increasing_terrace' or patternName == 'decreasing_terrace' else False
        isGorge = True if patternName == 'gorge' else False
        isPlainorPlateau = True if patternName == 'plain' or patternName == 'proper_plain' or patternName == 'plateau' or patternName == 'proper_plateau' else False
        isOnlySteady = True if patternName == 'steady' or patternName == 'steady_sequence' else False
        isValley = True if patternName == 'valley' else False
        if (isMaxFeature and (isTerrace or isGorge or isPlainorPlateau or isOnlySteady or isValley)):
            return
        ## Stop generation of 6 extra '_min_' functions
        isMinFeature = True if featureName == 'min' else False
        isSummitorPeak = True if patternName == 'peak' or patternName == 'summit' else False
        if (isMinFeature and isSummitorPeak):
            return
        ## Stop the generation of superfluous 'Range' functions, monotonic patterns needed only.
        isRange=True if featureName == 'range' else False
        isDecSeq=True if patternName == 'decreasing_sequence' or patternName == 'strictly_decreasing_sequence' else False
        isIncSeq=True if patternName == 'increasing_sequence' or patternName == 'strictly_increasing_sequence' else False
        isStrict=True if patternName == 'strictly_increasing_sequence' or patternName == 'strictly_decreasing_sequence' else False
        # Determine if pattern is monotonic or not. 
        # If you want to generate range functions with all patterns and not just monotonic ones comment out the 3 lines below.
        isMonotonic = True if isDec or isDecSeq or isInc or isIncSeq or isStrict else False
        if (isMonotonic != True and isRange) :
            return
        
        ## Proceed to generate the 253 functions
        self.writeComment(x+'_'+featureName+'_'+patternName+' : Exported Function')
        self.writeHeader(patternName, featureName, x)#, aggregatorName)
        
        letters=['C', 'D', 'R', 'H'] if isRange and (isDecSeq or isIncSeq or isStrict) else ['C', 'D', 'R'] #Initializes values
        for accumulator in ['C', 'D', 'R']:
            self.writeInitValue(accumulator, patternName, featureName, aggregatorName)

        if isRange and (isIncSeq):
            self.writeLine("H := math.Inf(1)")
        elif isRange and (isDecSeq):
            self.writeLine("H := math.Inf(-1)")
        
        self.writeEntryState(patternName)
        self.writeLine('DataLen := len(data)')
        self.writeLine('for i := 1; i < DataLen; i++ {')
        self.indent()
        self.indent()
        for accumulator in letters:
            self.writeLine(accumulator + 'temp := float64(' + accumulator+')')
        
        self.writeLine('if data[i] > data[i-1] {')
        if (isRange and isDecSeq):
            self.indent()
            self.writeLine('H = 0.0')
            self.dedent()
        elif (isRange and isIncSeq):
            self.indent()
            self.writeLine('H = data[i-1]')
            self.writeLine('H = math.Min(H, Htemp)')
            self.dedent()
        self.writeCore(patternName, featureName, aggregatorName, '<', isRange, isInc, isDec, isIncSeq, isDecSeq, isStrict)
        self.dedent()

        self.writeLine('} else if data[i] < data[i-1] {')##
        if (isRange and isDecSeq):
            self.indent()
            self.writeLine('H = data[i-1]')
            self.writeLine('H = math.Max(H, Htemp) // Holding onto the largest value for sequence')
            self.dedent()
        elif (isRange and isIncSeq):
            self.indent()
            self.writeLine("H = math.Inf(1)")
            self.dedent()
        self.writeCore(patternName, featureName, aggregatorName, '>', isRange, isInc, isDec, isIncSeq, isDecSeq, isStrict)
        self.dedent()

        self.writeLine('} else if data[i] == data[i-1] {')##
        if (isRange and isDecSeq and isStrict):
            self.indent()
            self.writeLine('H = 0.0')
            self.dedent()
        elif (isRange and isIncSeq and isStrict):
            self.indent()
            self.writeLine("H = math.Inf(1)")
            self.dedent()
        self.writeCore(patternName, featureName, aggregatorName, '=', isRange, isInc, isDec, isIncSeq, isDecSeq, isStrict)## Only for decreasing funcs rn
        self.dedent()

        self.writeLine('}')##
        self.dedent()
        self.writeLine('_ = Ctemp // Temporary fix')## Temporary fix for un-used variables
        self.writeLine('_ = Dtemp // Temporary fix')##
        self.writeLine('}')##
        self.dedent()
        self.writeLine('return ' + core.getValue(aggregatorName) + '(R, C)')
        self.dedent()
        self.writeLine('}')

    def writeCore(self, patternName, featureName, aggregatorName, sign, isRange, isInc, isDec, isIncSeq, isDecSeq, isStrict):
        self.indent()
        c = True
        for state in core.getPatternStates(patternName):
            if c :
                self.writeLine('if currentState == \'' + state + '\' {')
                c = False
            else:
                self.writeLine('} else if currentState == \'' + state + '\' {')
            self.indent()
            semantic = core.getNextSemantic(patternName, state, sign)

            if(isRange and (isIncSeq or isDecSeq or isStrict or isDec or isInc)):#Added
                for accumulator in ['C', 'D', 'H', 'R']:
                    update = core.getRangeUpdate(accumulator, semantic, patternName, featureName, aggregatorName)
                    if len(update) > 0:
                        self.writeLine(update)
            
            else:
                for accumulator in ['C', 'D', 'R']:
                    update = core.getUpdate(accumulator, semantic, patternName, featureName, aggregatorName)
                    if len(update) > 0:
                        self.writeLine(update)
            self.writeLine('currentState = \'' + core.getNextState(patternName, state, sign) + '\'')
            self.dedent()
        self.writeLine('}')
        

start_time = time.time()

c = CodeGeneratorBackend()

c.begin(tab="    ")

c.writeComment('----------------------------------------------------------------------------')
c.writeComment('This file was auto-generated on ' + datetime.now().strftime('%Y-%m-%d'))
c.writeComment('By Charles W. Jeffries.')
c.writeComment('Source Code : https://github.com/CharlieWJ/Time-Series-Patten-Mapping')
c.writeComment('----------------------------------------------------------------------------')
c.writeLine('')
c.writeLine('package generatedInGo')
c.writeLine('')
c.writeLine('import (')
c.writeLine('\t"math"')
c.writeLine(')')
c.writeLine('')
c.writeLine('func add(x float64, y float64) float64  { return (x + y) }')
c.writeLine('func diff(x float64, y float64) float64 { return math.Abs(x - y) } // The absolute difference')
c.writeLine('')

nb_func = 0 #Number of functions

#aggregator_names is used for the naming of the functions
#aggregators is used for obtaining the min/max using the "math" library.
for agg, agg_name in zip(aggregators,aggregator_names):
    for feature in features:
        for pattern in patterns:
            c.writeFunction(pattern, feature, agg, agg_name)
            nb_func = nb_func + 1

#my_file = open("D:\\TimeSeriesPatternMapping\\src\\generatedInGo\\generatedInGo.go", "w")
my_file = open("./generatedInGo/generatedInGo.go", "w")
my_file.write(c.end())
my_file.close()

exec_time = time.time() - start_time
nb_functions = len(aggregators) * len(features) * len(patterns) - (48+44+15+30+6) 
#48 corresponds to non monotonic range patterns, 44 corresponds to Max/Min_one funcs.
nb_func -= (48+44+15+30+6)

print('-----')
print('status : success')
print('functions generated : %d/%d' % (nb_func, nb_functions))
print("exec time : %s seconds" % exec_time)
print('-----')
