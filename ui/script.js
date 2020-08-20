var app = new Vue({
    el: '#app',
    data: {
        message: 'Hello Vue!'
    }
})



const Http = new XMLHttpRequest();
const url = 'http://localhost:9090';
Http.open("GET", url);
Http.send();

Http.onreadystatechange = (e) => {
    console.log(Http.responseText)
}