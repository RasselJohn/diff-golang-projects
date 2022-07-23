;(function () {
    // init form if it exists
    let form = document.getElementById('form');
    if (!form)
        return false;

    form.onsubmit = () => {
        let options = {method: 'POST', body: new FormData(form)};
        makeAjaxRequest(form.action, options);
        return false;
    };


    function openModal(fileName = "", errors = null) {
        // set modal
        let zipFile = document.getElementById('zipFile');
        if (fileName) {
            let fileUrl = `${document.location.origin}/user_files/${fileName}.zip`;
            zipFile.setAttribute('href', fileUrl);
            zipFile.setAttribute('style', "visibility: visible;");
        } else {
            zipFile.setAttribute('href', '#!');
            zipFile.setAttribute('style', "visibility: hidden;");
        }

        let modalIcon = document.getElementById('modalIcon');
        let modalMessageText = document.getElementById('modalMessageText');
        if (errors) {
            modalIcon.innerText = "report"
            modalMessageText.innerText = `Errors: ${errors.join("\n")}`
        } else {
            modalIcon.innerText = "thumb_up"
            modalMessageText.innerText = "Operation success!"
        }

        M.Modal.init(document.querySelector('.modal'), {}).open();
    }

    async function makeAjaxRequest(url, options) {
        let response = await fetch(url, options);

        if (response.status === 200 || response.status === 400) {
            let jsonResponse = await response.json();
            openModal(jsonResponse.file, jsonResponse.errors);
        } else {
            openModal("", ["Unknown error! Please, repeat later."]);
        }
    }
})();