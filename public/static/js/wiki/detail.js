const wikiBodyElement = document.getElementById("wiki-body");
const deleteButtonElement = document.getElementById("delete-wiki");
const currentWikiIdElement = document.getElementById("current_wiki_id");
const currentCategoryIdElement = document.getElementById("current_category_id");

const convertWikiBody = (event) => {
    wikiBodyElement.innerHTML = wikiBodyElement.innerText;
    wikiBodyElement.classList.remove("hidden");
}

const deleteWiki = async (event) => {
    event.preventDefault();
    const wikiId = currentWikiIdElement.value;
    const agreeConfirm = confirm("Are you sure that you want to delete this post?");
    if (agreeConfirm && wikiId) {
        try {
            const response = await fetch(`/wiki/${wikiId}`, {
                method: "DELETE"
            });
            if (response.ok) {
                let targetUrl = "/";
                if (!!currentCategoryIdElement.value) {
                    targetUrl = `/category/view/${currentCategoryIdElement.value}`;
                }
                window.location.href = targetUrl;
                return;
            }
            alert("Unknown issue, please try again!");
        } catch (e) {
            alert(e.message);
        }
    }
}

document.addEventListener("DOMContentLoaded", convertWikiBody);
deleteButtonElement.addEventListener("click", deleteWiki);