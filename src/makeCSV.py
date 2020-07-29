
def makeCSV(fileName):
    f = open(fileName, "r")
    dataHolder=[] #this will hold the data slices that have already been used to avoid duplicate csv file generation
    casenum=0
    case='case'
    for x in f:
        x=f.read(6)                         #determines if it is a method or not
        if(x=='method'):
            line=f.readline()
            length=len(line)-4
            for i in range(length):
                if line[i]=='{':
                    data=line[i+1:length:1]
                    if data in dataHolder:
                        if data == '3, 4, 2, 2, 5, 6, 6, 4, 4, 3, 1, 1, 4, 6, 4, 4': # This conditional was used to find which data slices corresponded to the correct test cases
                            print(line)
                        continue            #skips duplicate data
                    else:
                        dataHolder.append(data) #adds data slice
        else:
            continue
    f.close()
    # File creation
    for line in dataHolder:
        fl=open("../res/tests/"+case+str(casenum)+".csv",'w')
        length=(len(line))
        print(case+str(casenum)+': '+line)
        for i in range(length):
            if line[i]==',':
                continue
            elif line[i]==' ':
                fl.write('\n')
            else:
                fl.write(line[i])
        fl.close()
        casenum+=1

#print(makeCSV("inputs\\input_go_tests.txt"))

def makeCases(fileName):
    f = open(fileName, "r")
    names = open("names.go","a")
    for x in f:
        if x[0:6:1]=="Method":
            length=len(x)-4
            for i in range(length):
                if x[i]==']':
                    data=x[i+9:length:1]
                    if data == '':
                        temp=next(f)
                        written="\n{\n     "+x+"     "+temp+"},"
                        #print(written)
                        names.write(written)
    f.close()
    names.write('\n}')
    names.close()

#print(makeCases("formatting.txt"))