; (function () {
    // init form if it exists
    let form = document.getElementById('form');

    if (!form) return false;

    form.onsubmit = () => {
        let csrf = document.getElementsByName('csrfmiddlewaretoken')[0].value;
        let options = {
            method: 'POST',
            headers: { 'X-CSRFToken': csrf },
            body: new FormData(form)
        };
        makeAjaxRequest(form.action, options);
        return false;
    };


    function openModal(message = 'Неизвестная ошибка !') {
        // set modal message
        document.getElementById('modalMessage').innerText = message;
        M.Modal.init(document.querySelector('.modal'), {}).open();
    }

    function makeAjaxRequest(url, options) {
        fetch(url, options)
            .then((response) => response.json())
            .then((json) => {
                if (json.url) {
                    location = json.url;
                    return;
                }

                if (json.error) {
                    return openModal(json.error);
                } else {
                    return openModal();
                }
            })
            .catch(openModal);
    }
})();