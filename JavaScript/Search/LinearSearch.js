
function linearSearch(arr, key){
    for(let i = 0; i < arr.length; i++){
        if(arr[i] === key){
            return i
        }
    }
    return -1
}

let inputArr = [5, 2, 4, 6, 1, 3];

linearSearch(inputArr);
console.log;