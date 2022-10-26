import {printMsg1, printMsg2} from '../NumberGen/RandomInt.js'

function insertionSort (arr){
    let it = 0;
    for (let i = 1; i < arr.length; i++) {
      let j = i - 1
      
      let temp = arr[i]
      it++;
      while (j >= 0 && arr[j] > temp) {
        arr[j + 1] = arr[j]
        j--
        it++;
      }
      arr[j+1] = temp
    }
    process.stdout.write("\nINSERTION SORT\n");
    process.stdout.write("Iteraciones = "+it+"\n")
    return arr
  }

//Unsorted array
//TEST
// let arr = [234, 43, 55, 63, 5, 6, 235, 547];

// printMsg1(arr);
// printMsg2(insertionSort(arr));

export {insertionSort};