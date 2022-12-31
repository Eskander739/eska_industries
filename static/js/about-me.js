var changeData = 0;
var codeApprove = "";
var emailName = "";
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
            if (resp.email != undefined) {
                
                var name = document.getElementById("name");
                var email = document.getElementById("email");
                name.value  = `${resp.username}`;
                email.value  = `${resp.email}`;
                emailName = `${resp.email}`;
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
function cancelButton(){

    // var cancel = document.getElementById("cancel");
    var changeName = document.getElementById("change-name");
    var name = document.getElementById("name");

    name.setAttribute('disabled', 'disabled');
    changeName.innerHTML = "изменить";
    document.getElementsByClassName('input-block')[0].removeChild(cancel);
    changeName.onclick = changeButton;

}

function httpPost(){
    var xmlHttp = new XMLHttpRequest();

    

    var userName = document.getElementById("name").value;
    


    var url = `http://localhost:8000/account/about-me/change-name/${getCookie("EskaUser")}/${userName}`;
    xmlHttp.open( "POST", url, false ); // false for synchronous request
    xmlHttp.send( null );
    var cancel = document.getElementById("cancel");
    
    var changeName = document.getElementById("change-name");
    document.getElementById("name").setAttribute('disabled', 'disabled');

    changeName.innerHTML = "изменить";
    document.getElementsByClassName('input-block')[0].removeChild(cancel);
    changeName.onclick = changeButton;
    return xmlHttp.responseText;
}

function httpPostEmail(){

    var xmlHttp = new XMLHttpRequest();

    var url = `http://localhost:8000/account/about-me/change-email/${emailName}`;

    xmlHttp.onreadystatechange = function () {
        if (this.readyState == 4) {

        }

        if (this.status == 200) {
            var responses = JSON.parse(this.responseText);
            if (responses.code != undefined) {
                codeApprove = `${responses.code}`;
            }

            // alert(resp.email);
        }


        // end of state change: it can be after some time (async)

    };
    xmlHttp.open( "POST", url, false ); // false for synchronous request
    xmlHttp.send( null );
}

function httpPostApproveChangeEmail(){
    var changeEmail = document.getElementById("email");
    var approveCode = document.getElementById('code');

    if (approveCode.value == codeApprove) {
        if (changeEmail.value != ""){
            var xmlHttp = new XMLHttpRequest();
            var url = `http://localhost:8000/account/about-me/change-email/approve/${getCookie("EskaUser")}/${changeEmail.value}`;
            xmlHttp.open( "POST", url, false ); // false for synchronous request
            xmlHttp.send( null );
            var changeName = document.getElementById("change-email");
            var changeEmail = document.getElementById("email");
            var approveCode = document.getElementById('code');
            var approveButton = document.getElementById('approve');
            var impInfo = document.getElementById('impInfo');
            var impInfo2 = document.getElementById('impInfo2');
            changeName.innerHTML = "изменить";
            changeEmail.setAttribute('disabled', 'disabled');
            document.getElementsByClassName('input-block-email')[0].removeChild(approveCode);
            document.getElementsByClassName('input-block-email')[0].removeChild(approveButton);
            document.getElementsByClassName('input-block-email')[0].removeChild(impInfo);
            document.getElementsByClassName('input-block-email')[0].removeChild(impInfo2);


        }
    }

}

function changeButton(){

    var changeName = document.getElementById("change-name");
    // alert(changeName.innerHTML);
    if (changeName.innerHTML == "сохранить"){
        var name = document.getElementById("name");
    
        name.setAttribute('disabled', 'disabled');
        changeName.innerHTML = "изменить";
        document.getElementsByClassName('input-block')[0].removeChild(cancel);

    }

    if (changeName.innerHTML == "изменить"){

        var saveButton = document.createElement('button');
        saveButton.className = "cancel";
        saveButton.id = "cancel";
        saveButton.innerHTML = "отменить";
        saveButton.onclick = cancelButton;
        var name = document.getElementById("name");

        name.removeAttribute('disabled');
        changeName.innerHTML = "сохранить";
        changeName.onclick = httpPost;

        document.getElementsByClassName('input-block')[0].appendChild(saveButton);

    }



}


function changeButtonEmail(){
    var changeName = document.getElementById("change-email");

    if (changeName.innerHTML == "изменить"){
        changeData = 1;
        var changeEmail = document.getElementById("email");
        var approveCode = document.createElement('input');
        var ImpormantInfoBlock = document.createElement('div');
        var ImpormantInfoSecondBlock = document.createElement('div');
        var ImpormantInfo = document.createElement('label');
        var ImpormantInfoSecond = document.createElement('label');
        approveCode.className = "code";
        approveCode.id = "code";
        approveCode.placeholder = "Введите код отправленный на вашу почту";
        approveCode.className = "code-write";
        ImpormantInfoSecond.innerHTML = "Не перезагружайте страницу, при перезагрузке процесс смены почты прервется";
        ImpormantInfoSecond.style = "color:red!important;text-shadow: -1px 0 black, 0 1px black, 1px 0 black, 0 -1px black;font-family: arial black;"
        ImpormantInfoSecondBlock.style = "text-align: center;";
        ImpormantInfoSecondBlock.appendChild(ImpormantInfoSecond);
        ImpormantInfoSecondBlock.id = "impInfo2"
        ImpormantInfo.innerHTML = "Нажмите кнопку подтвердить когда введете новую почту и отправленный код";
        ImpormantInfoBlock.id = "impInfo";
        ImpormantInfoBlock.appendChild(ImpormantInfo);
        ImpormantInfoBlock.style = "text-align: center;";
        var approveButton = document.createElement('button');
        approveButton.className = "approve";
        approveButton.id = "approve";
        approveButton.innerHTML = "подтвердить";
        approveButton.onclick = httpPostApproveChangeEmail;
        changeName.innerHTML = "отменить";
        changeEmail.removeAttribute('disabled');
        changeEmail.value = "";
        changeEmail.placeholder = "Введите новую почту";
        // // saveButton.onclick = cancelButton;
        // var name = document.getElementById("name");

        // name.removeAttribute('disabled');
        // changeName.innerHTML = "сохранить";
        // changeName.onclick = httpPost;

        document.getElementsByClassName('input-block-email')[0].appendChild(approveCode);
        document.getElementsByClassName('input-block-email')[0].appendChild(approveButton);
        document.getElementsByClassName('input-block-email')[0].appendChild(ImpormantInfoBlock);
        document.getElementsByClassName('input-block-email')[0].appendChild(ImpormantInfoSecondBlock);
        httpPostEmail();

    }
    if (changeName.innerHTML == "отменить" && changeData == 0){
        var changeEmail = document.getElementById("email");
        var approveCode = document.getElementById('code');
        var approveButton = document.getElementById('approve');
        var impInfo = document.getElementById('impInfo');
        var impInfo2 = document.getElementById('impInfo2');
        document.getElementsByClassName('input-block-email')[0].removeChild(approveCode);
        document.getElementsByClassName('input-block-email')[0].removeChild(approveButton);
        document.getElementsByClassName('input-block-email')[0].removeChild(impInfo);
        document.getElementsByClassName('input-block-email')[0].removeChild(impInfo2);
        changeName.innerHTML = "изменить";
        changeEmail.value = emailName;
        changeEmail.setAttribute('disabled', 'disabled');

    }
    changeData = 0;





}


window.onload = function () {
    getUser();
}
