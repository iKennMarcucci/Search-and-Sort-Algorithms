//Get a random 5 digits number
function getRandomInt() {
    let num = Math.floor(Math.random() * 90000) + 10000;
    return num;
}

//Generate a 5 digits array with n lenght
function fillArr(n) {
    let arr = new Array(10);
    for (let i = 0; i < n; i++) {
        arr[i] = getRandomInt();

    }
    return arr;
}

console.log(fillArr(20));