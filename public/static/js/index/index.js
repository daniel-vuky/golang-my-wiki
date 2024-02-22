const listCategoryElement = document.getElementById("list-category");

const renderCategory = (listCategory) => {
    let listItemsElement = "";
    listCategory.map(function (category) {
        const name = category.name;
        let slicedName = name.slice(0, 45);
        if (name.length > 45) {
            slicedName += "...";
        }
        listItemsElement +=
            '<div class="list-items">\n' +
            '   <h1 class="title">\n' +
            '       <a href="/category/edit/' + category.category_id + '">' +
            '       <a href="/category/view/'+ category.category_id +'" title="' + name + '">' + slicedName + '</a> ' +
            '   </h1>\n' +
            '   <div class="time">Last Updated: ' + category.updated_at + '</div>\n' +
            '</div>';
    })
    // List item elements
    listCategoryElement.innerHTML = listItemsElement;
}

const renderError = (message) => {
    const errorElement = document.createElement("div");
    errorElement.classList.add("error-wrapper");
    errorElement.innerText = message
    listCategoryElement.innerHTML = errorElement;
}

const loadAllParentCategory = async (event) => {
    try {
        const response = await fetch("/category");
        if (response.ok) {
            const responseJson = await response.json();
            renderCategory(responseJson);
            return
        }
        renderError("Unknown issue, please try again!");
    } catch (e) {
        renderError(e.message);
    }
}

document.addEventListener("DOMContentLoaded", loadAllParentCategory);