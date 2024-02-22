const registerFormElement = document.getElementById("register-form");
const errorWrapper = document.getElementById("error-wrapper");

const emailElement = document.getElementById("email");
const passwordElement = document.getElementById("password");
const confirmPasswordElement = document.getElementById("confirm_password");
const submitBtnElement = document.getElementById("submit-btn");

const toggleFieldValidateErr = (fieldId, message, clear) => {
    if (!!clear) {
        const fieldElement = document.getElementById(fieldId);
        if (!!fieldElement) {
            fieldElement.classList.remove("field-error");
        }
        const fieldErrorElement = document.getElementById(`${fieldId}-error`);
        if (!!fieldErrorElement) {
            fieldErrorElement.remove();
        }
        submitBtnElement.disabled = false;
        return;
    }
    let newErrDiv = document.createElement("div");
    newErrDiv.id = `${fieldId}-error`;
    newErrDiv.classList.add("field-error-msg");
    newErrDiv.textContent = message;
    const fieldElement = document.getElementById(fieldId);
    fieldElement.parentNode.insertBefore(newErrDiv, fieldElement.nextSibling);
    fieldElement.classList.add("field-error");
    submitBtnElement.disabled = true;
}

const toggleShowErrorMessage = (msg) => {
    errorWrapper.textContent = msg;
    submitBtnElement.disabled = true;
}

const registerPost = async (event) => {
    event.preventDefault();
    toggleLoaderForBtn(
        {
            id: "submit-btn"
        },
        {
            tag_name: "div",
            class_list: "loader",
            id: "loader",
            text_content: ""
        },
    );
    let formFields = new URLSearchParams();
    for (const input of event.target.elements) {
        if (!!input.name) {
            formFields.append(input.name, input.value);
        }
    }
    try {
        const response = await fetch("/register", {
            method: "POST",
            credentials: "same-origin",
            headers: {
                "Content-type": "application/x-www-form-urlencoded"
            },
            body: formFields.toString()
        });
        if (response.ok) {
            window.location.href = "/";
            return;
        }
        const responseJson = await response.json();
        toggleShowErrorMessage(responseJson.message);
    } catch (error) {
        toggleShowErrorMessage("Unexpected error when try to create new account, please try again!");
    }
    toggleLoaderForBtn(
        {
            id: "loader"
        },
        {
            tag_name: "button",
            class_list: "submit-btn",
            id: "submit-btn",
            text_content: "Sign Up "
        },
    );
}

const asyncValidateEmail = async (event) => {
    try {
        toggleFieldValidateErr("email", "", true)
        const email = event.target.value;
        const url = `/user-existed?email=${encodeURI(email)}`;
        const response = await fetch(url);
        const responseJson = await response.json();
        if (responseJson["existed"]) {
            toggleFieldValidateErr(
                "email",
                "Email has been existed",
                false
            );
        }
    } catch (err) {
        console.error(err.message);
    }
}

const validateConfirmPassword = (event) => {
    toggleFieldValidateErr(
        "confirm_password",
        "",
        true
    );
    const confirmPassword = event.target.value;
    if (!!confirmPassword) {
        const password = passwordElement.value;
        if (password !== confirmPassword) {
            toggleFieldValidateErr(
                "confirm_password",
                "Confirm password need to be the same with password",
                false
            );
        }
    }
}

registerFormElement.addEventListener("submit", registerPost);
emailElement.addEventListener("change", asyncValidateEmail);
passwordElement.addEventListener("change", validateConfirmPassword);
confirmPasswordElement.addEventListener("change", validateConfirmPassword);