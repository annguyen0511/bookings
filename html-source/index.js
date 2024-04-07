let attention = Prompt();

// Example starter JavaScript for disabling form submissions if there are invalid fields
(() => {
    'use strict'

    window.addEventListener('load', () => {
        // Fetch all the forms we want to apply custom Bootstrap validation styles to
        const forms = document.querySelectorAll('.needs-validation')

        // Loop over them and prevent submission
        Array.from(forms).forEach(form => {
            form.addEventListener('submit', event => {
                if (!form.checkValidity()) {
                    event.preventDefault()
                    event.stopPropagation()
                }

                form.classList.add('was-validated')
            }, false)
        })
    }, false);
})()


document.getElementById("test-btn").addEventListener("click", () => {
    // notify("This is my message", "error");
    notifyMoDal("title", "This is my message", "success", "This is my message");
    attention.toast({ msg: "This is my message", icon:"error" });
})



function notify(msg, msgType) {
    notie.alert({
        type: msgType, // optional 'success', 'warning', 'error', 'info', 'neutral']
        text: msg,
    })
}


function notifyMoDal(title, text, icon, confirmButtonText) {
    Swal.fire({
        title: title,
        text: text,
        icon: icon,
        confirmButtonText: confirmButtonText,
    })
}

function Prompt() {

    let toast = c => {
        const {
            msg = "",
            icon = "success",
            position = "top-end",
        } = c;

        const Toast = Swal.mixin({
            toast: true,
            title: msg,
            position: position,
            icon: icon,
            showConfirmButton: false,
            timer: 3000,
            timerProgressBar: true,
            didOpen: (toast) => {
                toast.onmouseenter = Swal.stopTimer;
                toast.onmouseleave = Swal.resumeTimer;
            }
        });
        Toast.fire({});
    }
        let success = () => {
            console.log("clicked button and got success")
        }

        return {
            toast: toast,
            success: success,
        }


}