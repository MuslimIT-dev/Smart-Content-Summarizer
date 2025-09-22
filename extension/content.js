function getMainText() {
    let article = document.querySelector("article");
    if (article) return article.innerText.trim();

    let blocks = [...document.querySelectorAll("p, div")];
    let biggest = blocks.reduce((a, b) =>
      b.innerText.length > a.innerText.length ? b : a
    );
    return biggest.innerText.trim();
}