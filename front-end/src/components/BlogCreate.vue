<template>
  <div class="col-md-6 center">
    <form>
      <div class="form-group">
        <label for="exampleFormControlInput1">Blog Title</label>
        <input v-model="title" type="text" class="form-control" id="exampleFormControlInput1"
               placeholder="Enter blog title here">
      </div>
      <div class="form-group">
        <label for="exampleFormControlTextarea1">Blog Body</label>
        <textarea v-model="body" class="form-control" id="exampleFormControlTextarea1" rows="3"></textarea>
      </div>
      <button @click="createBlog" type="submit" class="btn btn-primary mb-2">Save</button>
    </form>
  </div>
</template>

<script>
  export default {
    name: "blog-create",
    data() {
      return {
        title: '',
        body: ''
      }
    },
    methods: {
      createBlog() {
        if (this.title && this.body) {
          fetch('http://localhost:3000/create-blog', {
            method: 'post',
            body: JSON.stringify({Title: this.title, Body: this.body})
          })
            .then((resp) => resp.json())
            .then((blog) => {
              this.$router.push('/');
            }).catch((err) => {
            this.$noty.error('Something wrong happened - ' + err)
          })
        } else {
          this.$noty.error('All fields are required')
        }
      }
    }
  }
</script>

<style scoped>
  .center {
    margin: 30px auto;
  }
</style>
