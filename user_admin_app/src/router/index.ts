import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'
import { useAuthStore } from '@/stores/auth'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView
    },
    {
      path: '/about',
      name: 'about',
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import('@/views/AboutView.vue'),
      meta: { requiresAuth: true }
    }
  ]
})

router.beforeEach((to, from, next) => {
  const authStore = useAuthStore()
  console.log(to, from, authStore.isLoggedIn)
  if (to.name === 'home' && authStore.isLoggedIn) {
    next({ name: 'about' })
  } else if (to.meta.requiresAuth && !authStore.isLoggedIn) {
    next({ name: 'home' })
  } else {
    next()
  }
})

export default router
