var globalIndex = 0;

async function getImage(imageId) {
    // alert(1);
    let xhr = new XMLHttpRequest();
    let url = "http://localhost:8000/get-image-by-id";

    xhr.onreadystatechange = function () {
        if (this.readyState == 4) {

        }

        if (this.status == 200) {
            var immgg = document.getElementById(imageId);

            var blob = new Blob([this.response], { type: "image/jpg" });
            var image = URL.createObjectURL(blob);
            immgg.src = image;
        }


        // end of state change: it can be after some time (async)

    };
    xhr.open("POST", url, true);
    xhr.responseType = "arraybuffer"; //ОБЯЗАТЕЛЬНЫЙ ТЕГ ЧТОБЫ ИЗОБРАЖЕНИЕ ЧИТАЛОСЬ!!!
    xhr.setRequestHeader("Content-Type", "application/json");
    var data = JSON.stringify({ "id": imageId });
    xhr.send(data);
    // alert(xhr.response);


}


async function getImage2(imageId) {

    let xhr = new XMLHttpRequest();
    let url = "http://localhost:8000/get-image-by-id";

    xhr.onreadystatechange = function () {
        if (this.readyState == 4) {

        }

        if (this.status == 200) {
            var imageMain = document.getElementById(`GlobalIndex_${imageId}`);
            var blob = new Blob([this.response], { type: "image/jpg" });
            var image = URL.createObjectURL(blob);
            imageMain.src = image;


        }


    };
    xhr.open("POST", url, true);
    xhr.responseType = "arraybuffer"; //ОБЯЗАТЕЛЬНЫЙ ТЕГ ЧТОБЫ ИЗОБРАЖЕНИЕ ЧИТАЛОСЬ!!!
    xhr.setRequestHeader("Content-Type", "application/json");
    var data = JSON.stringify({ "id": imageId });

    xhr.send(data);


}


function getData(callback) {
    var xhr = new XMLHttpRequest();
    // we defined the xhr
    var yourUrl = "http://localhost:8000/post-list";

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


    xhr.open('POST', yourUrl, true);
    xhr.send();

    // return data1;


}

function entryPosts(imageIndex) {

    var entryBig = document.createElement('div');
    entryBig.className = "entry";


    var entryContent = document.createElement('div');
    entryContent.className = "entry__content";

    var entryInfo = document.createElement('div');
    entryInfo.className = "entry__info";

    var bigImage = document.createElement('img');
    bigImage.id = imageIndex;

    var entryInfoLink = document.createElement('a');
    entryInfoLink.href = `http://localhost:8000/reading-page/${post.Id}`;
    entryInfoLink.appendChild(bigImage);

    var mainTitle = document.createElement('h1');
    var mainTitleLink = document.createElement('a');
    mainTitleLink.innerText = post.title;

    var DateInfo = document.createElement('ul');
    var DateInfo2 = document.createElement('li');
    DateInfo2.innerText = post.date;

    DateInfo.className = "entry__meta";
    DateInfo.appendChild(DateInfo2);

    mainTitle.appendChild(mainTitleLink);

    entryInfo.appendChild(entryInfoLink);
    entryInfo.appendChild(DateInfo);
    entryContent.appendChild(entryInfo);

    entryContent.appendChild(mainTitle);

    entryBig.appendChild(entryContent);

    return entryBig




}

function DataPosts() {
    let arr = new Array();
    var bigPostSmall = document.createElement('div');
    bigPostSmall.className = "featured__column featured__column--small";

    var resp = JSON.parse(this.responseText);
    for (index = 0; index < resp.posts.length; ++index) {
        // alert(1);
        //<a class="nav-link" href="http://localhost:8000/all_posts"><i></i>Cancel</a>
        post = resp.posts[index];
        // alert(1);
        var mainArticle = document.createElement("article");
        var aLink = document.createElement("a");
        var entryThumb = document.createElement("entry__thumb");

        entryThumb.className = "entry__thumb";
        aLink.className = "entry__thumb-link";
        aLink.href = `http://localhost:8000/reading-page/${post.Id}`;//${post.title}/${post.content}

        mainArticle.className = "masonry__brick entry format-standard";

        var postImage = document.createElement('img');
        postImage.id = post.Id;
        var entryText = document.createElement("div");
        var entryHeader = document.createElement("div");
        entryText.className = "entry__text";
        entryHeader.className = "entry__header";


        var entryTitle = document.createElement('h1');
        entryTitle.className = "entry__title";
        // entryTitle.innerText = post.title;
        entryTitle.href = `http://localhost:8000/reading-page/${post.Id}`;//${post.title}/${post.content};
        var H1Link = document.createElement('a');
        H1Link.href = `http://localhost:8000/reading-page/${post.Id}`;//${post.title}/${post.content};
        entryTitle.innerText = post.title;
        H1Link.appendChild(entryTitle);

        var entryDate = document.createElement('div');
        entryDate.className = "entry__date";

        var entryDate2 = document.createElement('a');
        entryDate2.href = `http://localhost:8000/reading-page/${post.Id}`;//${post.title}/${post.content};
        entryDate2.innerText = post.date;
        entryDate.appendChild(entryDate2);


        var entryExcerpt = document.createElement('div');
        var textMain = document.createElement('a');
        textMain.href = `http://localhost:8000/reading-page/${post.Id}`;//${post.title}/${post.content}
        textMain.innerText = post.content;
        textMain.style = "color: black; overflow: hidden; display: -webkit-box; -webkit-line-clamp: 3; -webkit-box-orient: vertical;";
        entryTitle.href = `http://localhost:8000/reading-page/${post.Id}`;//${post.title}/${post.content};
        entryTitle.style = "color: black";

        entryExcerpt.className = "entry__excerpt";

        // mainArticle.style = "fade-up";
        entryHeader.appendChild(entryDate);
        entryText.appendChild(entryHeader);
        entryText.appendChild(entryExcerpt);
        entryHeader.appendChild(H1Link);
        aLink.appendChild(postImage);
        entryThumb.appendChild(aLink);
        entryExcerpt.appendChild(textMain);
        mainArticle.appendChild(entryThumb);
        mainArticle.appendChild(entryText);

        document.getElementsByClassName('masonry')[0].appendChild(mainArticle);

        if (index == 0) {
            var bigPost = document.createElement('div');
            bigPost.className = "featured__column featured__column--big";
            var data = entryPosts(`GlobalIndex_${post.Id}`)
            arr.push(post.Id);
            bigPost.appendChild(data);
            // alert(globalIndex);
            document.getElementsByClassName('featured')[0].appendChild(bigPost);

        }

        if (index == 1 || index == 2) {

            if (index == 1) {
                arr.push(post.Id);
                var data = entryPosts(`GlobalIndex_${post.Id}`)

            }
            if (index == 2) {
                arr.push(post.Id);
                var data = entryPosts(`GlobalIndex_${post.Id}`)
            }
            globalIndex = index;
            // alert(globalIndex);
            bigPostSmall.appendChild(data);
        }


        if (index == 2) {
            document.getElementsByClassName('featured')[0].appendChild(bigPostSmall);

        }


        //     var last_id = main_div.children.length;
        //     var last_elem = window.getComputedStyle(main_div.children[last_id - 1]).left;
        //     var h1Data = document.createElement('h1');
        //     h1Data.className = "zagolovok";
        //     var title = document.createElement('label');

        //     var idd = document.createElement('label');
        //     idd.className = "idForAdd";
        //     idd.innerText = post.Id;
        //     idd.style = "display: none";
        //     postImage.id = post.Id;
        //     postImage.style = "display: block;width: auto;height: auto;max-width: 100%;";



        //     h1Data.innerText = post.title;
        //     title.className = 'title';
        //     title.appendChild(h1Data);
        //     var content = document.createElement('label');
        //     content.className = "content";
        //     content.innerText = post.content;
        //     content.style = "overflow: hidden;display: -webkit-box;-webkit-line-clamp: 3;-webkit-box-orient: vertical"
        //     var menuItem = document.createElement('div');
        //     menuItem.className = 'menu--item';
        //     var iDiv = document.createElement('figure');

        //     // iDiv.style.left = `300px`;
        //     // preee.appendChild(content);
        //     menuItem.style = `


        //     box-shadow:
        //      0 1px 4px rgba(0, 0, 0, .3),
        //      -23px 0 20px -23px rgba(0, 0, 0, .8),
        //      23px 0 20px -23px rgba(0, 0, 0, .8),
        //      0 0 40px rgba(0, 0, 0, .1) inset;
        //   }`;
        //     menuItem.style.left = `${Number(last_elem.replace("px", "")) + 460}px`;

        //     iDiv.appendChild(title);
        //     iDiv.appendChild(buttonRemove);
        //     iDiv.appendChild(buttonEdit);
        //     iDiv.appendChild(content);
        //     iDiv.appendChild(postImage);
        //     iDiv.appendChild(idd);
        //     // iDiv.id = 'block';


        //     menuItem.appendChild(iDiv);

        //     document.getElementsByClassName('menu--wrapper')[0].appendChild(menuItem);
        getImage(post.Id, globalIndex);
    }
    getImage2(arr[0]);
    getImage2(arr[1]);
    getImage2(arr[2]);

}

function getCookie(name) {
    const value = `; ${document.cookie}`;
    const parts = value.split(`; ${name}=`);
    if (parts.length === 2) return parts.pop().split(';').shift();
}

async function getUser() {
    // alert(1);
    let xhr = new XMLHttpRequest();
    let url = `http://localhost:8000/user/${getCookie("EskaUser")}`;

    xhr.onreadystatechange = function () {
        if (this.readyState == 4) {

        }

        if (this.status == 200) {
            var resp = JSON.parse(this.responseText);
            if (resp.email != undefined){
                var username = document.getElementById("userName");
                var immgg = document.getElementById("userLogo");
                var link = document.getElementById("userLink");
                immgg.src = "img/user-registered.png"
                username.innerText = resp.username;
                
                link.href = "http://localhost:8000/account/about-me";
            }

            // alert(resp.email);
        }


        // end of state change: it can be after some time (async)

    };
    xhr.open("POST", url, true);
    xhr.setRequestHeader("Content-Type", "application/json");
    xhr.send();
    // alert(xhr.response);


}

window.onload = function () {
    // setCookie('ppkcookie','testcookie',7);
    // alert(document.cookie);
    // alert(getCookie("test"))
    getData(DataPosts);
    getUser();
}

// function addPost() {
//     getData(DataPosts);




//     // alert(window.getComputedStyle(main_div.children[1]).left);
//     // alert(main_div.children[last_id-1]);


// }