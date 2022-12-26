async function getImage() {
    // alert(1);
    let xhr = new XMLHttpRequest();
    let url = "http://localhost:8000/get-image-by-id";

    xhr.onreadystatechange = function () {
        if (this.readyState == 4) {

        }

        if (this.status == 200) {
       
            var immgg = document.getElementsByClassName("mainImage")[0];
            var blob = new Blob([this.response], { type: "image/jpg" });
            immgg.src = URL.createObjectURL(blob);

        }


        // end of state change: it can be after some time (async)

    };
    var labelId = document.getElementById("IdReadPost");
    xhr.open("POST", url, true);

    xhr.responseType = "arraybuffer"; //ОБЯЗАТЕЛЬНЫЙ ТЕГ ЧТОБЫ ИЗОБРАЖЕНИЕ ЧИТАЛОСЬ!!!
    xhr.setRequestHeader("Content-Type", "application/json");
    var data = JSON.stringify({ "id": labelId.innerText });
    xhr.send(data);
    // alert(xhr.response);


}

function getData(callback) {
    var xhr = new XMLHttpRequest();
    // we defined the xhr
    var yourUrl = "http://localhost:8000/get-post-by-id";

    xhr.onreadystatechange = function () {
        if (this.readyState == 4) {
     
            callback.apply(xhr);

        }

        if (this.status == 200) {
            var data = JSON.parse(this.responseText);



            // we get the returned data
        }


        // end of state change: it can be after some time (async)

    };
    var labelId = document.getElementById("IdReadPost");
    // alert(labelId.innerText);

    // xhr.setRequestHeader("Content-Type", "application/json");

    var data = JSON.stringify({ "id": labelId.innerText });

    xhr.open('POST', yourUrl, true);
    xhr.send(data);

    // return data1;


}
function DataPosts() {
 
    var post = JSON.parse(this.responseText);
    // alert(post);
    var mainPost = document.createElement("div");
    mainPost.className = "post-type";

    var mainImage = document.createElement("img");
    mainImage.className = "mainImage";

    var title = document.createElement("h1");
    title.className = "title";
    title.innerText = post.title

    var content = document.createElement("content");
    content.className = "content";
    content.innerText = post.content

    mainPost.appendChild(mainImage);
    mainPost.appendChild(title);
    mainPost.appendChild(content);
    document.getElementsByClassName('body')[0].appendChild(mainPost);
    getImage();
}

window.onload = function () {
    
    getData(DataPosts);
}

// function addPost() {
//     getData(DataPosts);




//     // alert(window.getComputedStyle(main_div.children[1]).left);
//     // alert(main_div.children[last_id-1]);


// }