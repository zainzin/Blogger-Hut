import Vue from 'vue'
import Router from 'vue-router'
import Main from '@/components/Main'
import Authors from '@/components/Authors'
import Author from '@/components/Author'
import Blogs from '@/components/Blogs'
import BlogCreate from '@/components/BlogCreate'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Main',
      component: Main
    },
    {
      path: '/authors',
      name: 'Authors',
      component: Authors
    },
    {
      path: '/author',
      name: 'Author',
      component: Author
    },
    {
      path: '/blogs',
      name: 'Blogs',
      component: Blogs
    },
    {
      path: '/create-blog',
      name: 'BlogCreate',
      component: BlogCreate
    }
  ]
})
