const currentCategoryIdElement = document.getElementById("current_category_id");
const shortDescriptionElement = document.getElementById("short-description");
const listWikiElement = document.getElementById("list-wiki");
const deleteCategoryElement = document.getElementById("delete-category");

const getCurrentCategoryId = () => {
    return currentCategoryIdElement.value;
}

const deleteCategory = async (event) => {
    event.preventDefault();
    const categoryId = getCurrentCategoryId();
    const agreeDelete = confirm("Are you sure that you want to delete this category");
    if (agreeDelete && categoryId) {
        try {
            const response = await fetch(`/category/${categoryId}`, {method: "DELETE"});
            if (response.ok) {
                window.location.href = "/";
                return;
            }
            alert("Unknown issue, please try again!")
        } catch (e) {
            alert(e.message);
        }
    }
}

const renderWiki = (listWiki) => {
    let listItemsElement = "";
    listWiki.map(function (wiki) {
        const title = wiki.title;
        let slicedTitle = title.slice(0, 45);
        if (title.length > 45) {
            slicedTitle += "...";
        }
        listItemsElement +=
            '<div class="list-items">\n' +
            '   <h1 class="title">\n' +
            '       <a href="/wiki/view/'+ wiki.wiki_id +'" title="' + name + '">' + slicedTitle + '</a> ' +
            '   </h1>\n' +
            '   <div class="time">Last Updated: ' + wiki.updated_at + '</div>\n' +
            '</div>';
    })
    // List item elements
    listWikiElement.innerHTML = listItemsElement;
}

const renderError = (message) => {
    const errorElement = document.createElement("div");
    errorElement.classList.add("error-wrapper");
    errorElement.innerText = message
    listWikiElement.innerHTML = errorElement;
}

const loadSubCategory = async (event) => {
    try {
        const currentCategoryId = getCurrentCategoryId();
        if (!currentCategoryId) {
            return;
        }
        const response = await fetch(`/wiki?category_id=${currentCategoryId}`);
        if (response.ok) {
            const responseJson = await response.json();
            renderWiki(responseJson);
            if (responseJson.length === 0) {
                deleteCategoryElement.parentElement.classList.remove("hidden");
            }
            return
        }
        renderError("Unknown issue, please try again!");
    } catch (e) {
        renderError(e.message);
    }
}

shortDescriptionElement.innerHTML = shortDescriptionElement.innerText;

document.addEventListener("DOMContentLoaded", loadSubCategory);
deleteCategoryElement.addEventListener("click", deleteCategory)