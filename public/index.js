const bookList = document.getElementById('books');
const form = document.getElementById('form');

const URL = '/books';

window.onload = async () => {
  const res = await fetch(URL);
  const books = await res.json();

  books.forEach(addBookUI);
}

form.onsubmit = async (e) => {
  e.preventDefault();
  const title = form.title.value;
  const author = form.author.value;

  if(title === '' || author === '') {
    return;
  }

  const res = await fetch(URL, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      title,
      author
    })
  });
  const newBook = await res.json();
  addBookUI(newBook);
  form.reset();
}

function addBookUI(book) {
  const li = document.createElement('li');
    const a = document.createElement('a');

    a.href = `./book.html?bookid=${book.id}`
    a.innerText = book.title;
    li.appendChild(a)
    bookList.appendChild(li);
}
