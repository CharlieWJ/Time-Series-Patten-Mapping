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
        self.writeHeader(patternName, featureName, x)#, aggregatorName)
        for accumulator in ['C', 'D', 'R']:
            self.writeInitValue(accumulator, patternName, featureName, aggregatorName)
        self.writeEntryState(patternName)
        #self.writeLine('for i in range(1,len(data)):') #'range' was originally 'xrange', was changed because Python 3 no longer supports xrange
        self.writeLine('for i := 1; i < len(data); i++ {')
        self.indent()
        #self.writeLine('if(i < len(data)):')
        self.writeLine('if i<len(data) {')
        self.indent()
        for accumulator in ['C', 'D', 'R']:
            #self.writeLine(accumulator + '_temp = ' + accumulator)
            self.writeLine(accumulator + '_temp := ' + accumulator)
#        self.writeLine('if data[i] > data[i-1]:')
        self.writeLine('if data[i] > data[i-1] {')
        self.writeCore(patternName, featureName, aggregatorName, '<')
        self.dedent()
        self.writeLine('} else if data[i] < data[i-1] {')##
        #self.writeLine('elif data[i] < data[i-1]:')
        #self.writeLine('else if data[i] < data[i-1] {')
        self.writeCore(patternName, featureName, aggregatorName, '>')
        self.dedent()
        self.writeLine('} else if data[i] == data[i-1] {')##
        #self.writeLine('elif data[i] == data[i-1]:')
        #self.writeLine('else if data[i] == data[i-1] {')
        self.writeCore(patternName, featureName, aggregatorName, '=')
        self.dedent()
        self.writeLine('}')##
        self.dedent()
        self.writeLine('}')##
        self.dedent()
        self.writeLine('}')##
        self.writeLine('return ' + core.getValue(aggregatorName) + '(R,C)')
        self.dedent()
#        self.writeLine('')
        self.writeLine('}')

    def writeCore(self, patternName, featureName, aggregatorName, sign):
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
c.writeComment('By Charles W. Jeffries using an altered Python Script.')
c.writeComment('The original script was provided by: Florine Cercle & Denis Allard')
c.writeComment('Original Source Code : https://github.com/allarddenis/time-series-pattern-recognition')
c.writeComment('----------------------------------------------------------------------------')
c.writeLine('')
c.writeLine('package generatedInGo')
c.writeLine('')
c.writeLine('import(')
c.writeLine('\t"fmt"')
c.writeLine('\t"math"')
c.writeLine(')')
c.writeLine('')
c.writeLine('func add(x float64, y float64) float64 { return (x+y) }')
c.writeLine('type Tuple struct { a,b interface{} }')
c.writeLine('')
c.writeLine('func GetMax(x float64, y float64) float64 {')
c.writeLine('    if x>=y { ')
c.writeLine('        return x')
c.writeLine('    } else { return y }')
c.writeLine('}')
c.writeComment('Currently not using this method')
c.writeLine('') 

c.writeLine('func GetMin(x float64, y float64) float64 {') 
c.writeLine('    if x>=y { ')
c.writeLine('        return y')
c.writeLine('    } else { return x }')
c.writeLine('}')
c.writeComment('Currently not using this method')
c.writeLine('')
c.writeLine('')

nb_func = 0 #Number of functions

#aggregator_names is used for the naming of the functions
#aggregators is used for obtaining the min/max using the "math" library.
for agg, agg_name in zip(aggregators,aggregator_names):
    for feature in features:
        for pattern in patterns:
            c.writeFunction(pattern, feature, agg, agg_name)
            nb_func = nb_func + 1

#my_file = open("D:\\TimeSeriesPatternMapping\\src\\generated\\generated.py", "w") #Change the path depending on where it is located on your machine
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

### TODO: have an interface, and have a get max/min function for that interface if possible.