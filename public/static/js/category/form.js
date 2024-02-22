const categoryFormElement = document.getElementById("category-form");

const createCategoryPost = async (event) => {
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
    let formFields = {};
    for (let input of event.target.elements) {
        if (!!input.name) {
            formFields[input.name] = input.value;
        }
    }
    try {
        let options = {
            method: "POST",
            credentials: "same-origin",
            headers: {
                "Content-type": "application/json"
            },
            body: JSON.stringify(formFields)
        }
        let uri = "/category";
        if (formFields.hasOwnProperty("id") && formFields["id"] > 0) {
            options["method"] = "PUT";
            uri = `${uri}/${formFields["id"]}`;
        }
        const response = await fetch(uri, options);
        if (response.ok) {
            let targetUrl = "/";
            if (formFields.hasOwnProperty("id") && formFields["id"] > 0) {
                targetUrl = `${window.location.origin}/category/view/${formFields["id"]}`;
            }
            window.location.href = targetUrl;
            return;
        }
        alert("Unknown issue, please try again!");
    } catch (e) {
        alert(e.message);
    }
    toggleLoaderForBtn(
        {
            id: "loader"
        },
        {
            tag_name: "button",
            class_list: "submit-btn",
            id: "submit-btn",
            text_content: "Submit"
        },
    );
}

categoryFormElement.addEventListener("submit", createCategoryPost);