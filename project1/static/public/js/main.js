;(function () {
    // init form if it exists
    let form = document.getElementById('form');

    if (!form) return false;

    form.onsubmit = () => {
        // let csrf = document.getElementsByName('csrfmiddlewaretoken')[0].value;
        let options = {
            method: 'POST',
            // headers: { 'X-CSRFToken': csrf },
            body: new FormData(form)
        };
        makeAjaxRequest(form.action, options);
        return false;
    };


    function openModal(fileName, errors) {
        // set modal message
        document.getElementById('modalMessage').innerText = message;
        M.Modal.init(document.querySelector('.modal'), {}).open();
    }

    async function makeAjaxRequest(url, options) {
        let response = await fetch(url, options)
        if (response.status == 200) {
            let json = await response.json();
        } else {
            openModal(json.error);
        }


    }
})();