import {printMsg1, printMsg2} from '../NumberGen/RandomInt.js'

function selectionSort(arr) { 
    let n = arr.length;
    let it = 0;
    for(let i = 0; i < n; i++) {
        // Finding the smallest number in the subarray
        let min = i;
        for(let j = i+1; j < n; j++){
            if(arr[j] < arr[min]) {
                min=j; 
            }
         }
         if (min != i) {
             // Swapping the elements
             let tmp = arr[i]; 
             arr[i] = arr[min];
             arr[min] = tmp;      
        }
        it++;
    }
    process.stdout.write("\nSELECTION SORT\n");
    process.stdout.write("Iteraciones = "+it+"\n")
    return arr;
}

//Unsorted array
//TEST
// let arr = [234, 43, 55, 63, 5, 6, 235, 547];

// printMsg1(arr);
// printMsg2(selectionSort(arr));

export {selectionSort};