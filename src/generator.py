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
        self.writeLine('func ' + aggregatorName + '_' + featureName + '_' + patternName + '(data []float64) float64{')
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
        self.writeComment(x+'_'+featureName+'_'+patternName+' : Exported Function')
        self.writeHeader(patternName, featureName, x)#, aggregatorName)
        for accumulator in ['C', 'D', 'R']:
            self.writeInitValue(accumulator, patternName, featureName, aggregatorName)
        
        isRangeSeq=False                #Currently only Decreasing
        isRangeStrict=False             #Detect if func is range_strictly_decreasing_sequence
        if (featureName == 'range' and (patternName == 'decreasing_sequence')):# or patternName == 'increasing_sequence')):
            isRangeSeq=True
            self.writeLine("H := 0.0")
        
        if (featureName == 'range' and (patternName == 'strictly_decreasing_sequence')):# or patternName == 'strictly_increasing_sequence')):
            isRangeStrict=True
            self.writeLine("H := 0.0")
        
        self.writeEntryState(patternName)
        self.writeLine('for i := 1; i < len(data); i++ {')
        self.indent()
        self.writeLine('if i < len(data) {')
        self.indent()
        for accumulator in ['C', 'D', 'R']:
            self.writeLine(accumulator + 'temp := float64(' + accumulator+')')
        if (isRangeSeq or isRangeStrict):
            self.writeLine('Htemp := float64(H)')
        
        self.writeLine('if data[i] > data[i-1] {')
        if (isRangeSeq or isRangeStrict):
            self.indent()
            self.writeLine('H = 0.0')
            self.dedent()
        self.writeCore(patternName, featureName, aggregatorName, '<', isRangeSeq, isRangeStrict)
        self.dedent()

        self.writeLine('} else if data[i] < data[i-1] {')##
        if (isRangeSeq or isRangeStrict):
            self.indent()
            self.writeLine('H = data[i-1]')
            self.writeLine('H = math.Max(H, Htemp) // Holding onto the largest value for sequence')
            self.dedent()
        self.writeCore(patternName, featureName, aggregatorName, '>', isRangeSeq, isRangeStrict)
        self.dedent()

        self.writeLine('} else if data[i] == data[i-1] {')##
        if (isRangeStrict):
            self.indent()
            self.writeLine('H = 0.0')
            self.dedent()
        self.writeCore(patternName, featureName, aggregatorName, '=', isRangeSeq, isRangeStrict)
        self.dedent()

        self.writeLine('}')##
        self.dedent()
        self.writeLine('_ = Ctemp')## Temporary fix for un-used variables
        self.writeLine('_ = Dtemp')##
        self.writeLine('_ = Rtemp')##
        self.writeLine('}')##
        self.dedent()
        self.writeLine('}')##
        self.writeLine('return ' + core.getValue(aggregatorName) + '(R, C)')
        self.dedent()
        self.writeLine('}')

    def writeCore(self, patternName, featureName, aggregatorName, sign, isRangeSeq, isRangeStrict):
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

            if(isRangeSeq or isRangeStrict):#Added
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
c.writeComment('The original script was provided by: Florine Cercle & Denis Allard')
c.writeComment('Original Source Code : https://github.com/allarddenis/time-series-pattern-recognition')
c.writeComment('----------------------------------------------------------------------------')
c.writeLine('')
c.writeLine('package generatedingo')
c.writeLine('')
c.writeLine('import(')
c.writeLine('\t"math"')
c.writeLine(')')
c.writeLine('')
c.writeLine('func add(x float64, y float64) float64 { return (x+y) }')
c.writeLine('func diff(x float64, y float64) float64 { return (math.Abs(y-x)) }')
c.writeLine('')

nb_func = 0 #Number of functions

#aggregator_names is used for the naming of the functions
#aggregators is used for obtaining the min/max using the "math" library.
for agg, agg_name in zip(aggregators,aggregator_names):
    for feature in features:
        for pattern in patterns:
            c.writeFunction(pattern, feature, agg, agg_name)
            nb_func = nb_func + 1

my_file = open("D:\\TimeSeriesPatternMapping\\src\\generatedInGo\\generatedInGo.go", "w")
my_file.write(c.end())
my_file.close()

exec_time = time.time() - start_time
nb_functions = len(aggregators) * len(features) * len(patterns)

print('-----')
print('status : success')
print('functions generated : %d/%d' % (nb_func, nb_functions))
print("exec time : %s seconds" % exec_time)
print('-----')
