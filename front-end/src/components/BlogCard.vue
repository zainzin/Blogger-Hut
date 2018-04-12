<template>
  <div class="card" style="width: 18rem; margin: 5px;">
    <img class="card-img-top" v-bind:src="'https://picsum.photos/200/200/?image='+blog.ID * (Math.round(Math.random() * 10))" alt="Card image cap">
    <div class="card-body">
      <h5 class="card-title">{{blog.Title}}</h5>
      <p class="card-text">{{blog.Body.substring(0,20)}}</p>
      <p class="card-subtitle">Created At: {{createdAt}}</p>
      <p class="card-subtitle">Last Updated: {{updatedAt}}</p>
      <a href="#" class="btn btn-primary">View</a>
    </div>
  </div>
</template>

<script>
    import moment from 'moment'
    export default {
        name: "blog-card",
      data(){
          return {
            createdAt: moment(this.blog.CreatedAt).format("MM-DD-YYYY HH:mm"),
            updatedAt: moment(this.blog.UpdatedAt).format("MM-DD-YYYY HH:mm")
          }
      },
      props: ["blog"],
    }
</script>

<style scoped>

</style>
