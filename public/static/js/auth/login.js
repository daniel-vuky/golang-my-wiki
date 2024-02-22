const loginFormElement = document.getElementById("login-form");
const errorWrapper = document.getElementById("error-wrapper");

const toggleShowErrorMessage = (msg) => {
    errorWrapper.textContent = msg;
}

const loginPost = async (event) => {
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
        const response = await fetch("/login", {
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
        toggleShowErrorMessage("Unexpected error when try to login, please try again!");
    }
    toggleLoaderForBtn(
        {
            id: "loader"
        },
        {
            tag_name: "button",
            class_list: "submit-btn",
            id: "submit-btn",
            text_content: "Login "
        },
    );
}

loginFormElement.addEventListener("submit", loginPost);