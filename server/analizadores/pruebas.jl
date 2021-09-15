function swap(i, j, arr)
    temp = arr[i];
    arr[i] = arr[j];
    arr[j] = temp;
    return arr;
end;

function bubbleSort(arr)
    for i in 0:(length(arr) - 1)
        for j in 1:(length(arr) - 1)
            if(arr[j] > arr[j + 1])
                arr =swap(j, j+1, arr);
            end;
        end;
    end;
               return arr;
end;

function insertionSort(arr)

    for i in 1:length(arr)
        j = i;
        temp = arr[i];
        while(j > 1 && arr[j - 1] > temp)
            arr[j] = arr[j-1];
            j = j - 1;
        end;
        arr[j] = temp;
    end;
    return arr;
end;

arreglo = [32,7*3,7,89,56,909,109,2,9,9874^0,44,3,820*10,11,8*0+8,10];
arreglo = bubbleSort(arreglo);
print("BubbleSort => ",arreglo);
println("");

arregloo = [32,7*3,7,89,56,909,109,2,9,9874^0,44,3,820*10,11,8*0+8,10];
arregloo[1] = arregloo[2] - 21+4;
arregloo = insertionSort(arregloo);
print("InsertionSort => ",arregloo);