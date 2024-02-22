const actionButtons = document.getElementById("action-buttons");

const toggleLoaderForBtn = (elementToRemove, elementToReplace) => {
    actionButtons.removeChild(document.getElementById(elementToRemove.id));
    let newBtn= document.createElement(elementToReplace["tag_name"]);
    newBtn.classList.add(elementToReplace["class_list"]);
    newBtn.id = elementToReplace["id"];
    newBtn.textContent = elementToReplace["text_content"];
    actionButtons.append(newBtn);
}