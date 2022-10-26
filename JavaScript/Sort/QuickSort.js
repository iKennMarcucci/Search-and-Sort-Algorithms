import { printMsg1, printMsg2 } from '../NumberGen/RandomInt.js'

function quickSort(arr, start = 0, end = arr.length + 1) {
  let pivot = arr[start];
  let swapIndex = start;
  let it = 0;

  for (let i = start + 1; i < arr.length; i++) {
    if (pivot > arr[i]) {
      swapIndex++;
      swap(arr, i, swapIndex);
    }
    it++;
  }
  swap(arr, start, swapIndex);

  process.stdout.write("\nQUICK SORT\n");
  process.stdout.write("Iteraciones = "+it+"\n")
  
  return swapIndex;
  
  function swap(arr, firstIndex, secondIndex) {
    [arr[firstIndex], arr[secondIndex]] = [arr[secondIndex], arr[firstIndex]];
  }
}

// Unsorted array
// let arr = [234, 43, 55, 63, 5, 6, 235, 547];
// let high = (arr.length)-1;

// printMsg1(arr);
// quickSort(arr)
// printMsg2(arr);

export {quickSort};