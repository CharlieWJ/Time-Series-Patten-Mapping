
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
                        #if data == '7, 6, 5, 6, 5, 4, 1, 4, 7, 5, 4, 2, 5, 4, 3, 3': # This conditional was used to find which data slices corresponded to the correct test cases
                        #    print(line)
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



print(makeCSV("inputs\\input_go_tests.txt")) 
    