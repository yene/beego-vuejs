<template>
  <div class="hello">
    {{errorMessage}}
    <div v-show="loaded">
      <h3>Free books</h3>
      <ul>
        <li v-for="(book) in books" :key="book.isbn">
          <a :href="'https://www.goodreads.com/search?utf8=%E2%9C%93&query=' + book.isbn" target="_blank" rel="noopener">{{book.title}}</a>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  props: {

  },
  data() {
    return {
      loaded: false,
      errorMessage: '',
      books: [],
    }
  },
  beforeCreate() {
    // TODO: timeout, handle 404, handle body.errorCode != 0
    fetch('/books/all.json', {
      headers: {
        'Accept': 'application/json',
      },
    })
    .then(response => response.json())
    .then((data) => {
      if (data.errorCode != 0) {
        this.errorMessage = data.error
        return
      }
      this.books = data.books
      this.loaded = true
    })
    .catch((err) => {
      this.errorMessage = err
    });

  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  /*display: inline-block;*/
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
