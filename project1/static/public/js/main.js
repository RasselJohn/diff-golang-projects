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
            zipFile.hidden = false;
        } else {
            zipFile.setAttribute('href', '#!');
            zipFile.hidden = true;
        }

        let modalMessage = document.getElementById('modalMessage');
        if (errors) {
            modalMessage.innerText = `Errors: ${errors.join("\n")}`
            modalMessage.hidden = false
        } else {
            modalMessage.innerText = ""
            modalMessage.hidden = true
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