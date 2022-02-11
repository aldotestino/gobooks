const bookTitle = document.getElementById('title');
const bookAuthor = document.getElementById('author');
const bookAuthorImage = document.getElementById('pic')

const URL = '/book';
const avatarUrl = (author) => `https://avatars.dicebear.com/api/croodles-neutral/${author}.svg`


window.onload = async () => {
  const queryString = window.location.search;
  const urlParams = new URLSearchParams(queryString);
  const bookid = urlParams.get('bookid')

  const res = await fetch(`${URL}/${bookid}`);
  const book = await res.json();
  console.log(book);

  document.title = `${book.title} - ${book.author}`;
  bookTitle.innerText = book.title;
  bookAuthor.innerText = book.author
  bookAuthorImage.src = avatarUrl(book.author)
}
