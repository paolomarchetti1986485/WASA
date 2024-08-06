import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import ProfileView from '@/views/ProfileView.vue'


const routes = [
  { path: '/login', component: LoginView },
  { 
    path: '/', 
    component: HomeView,
    meta: { requiresAuth: true } // Aggiungi meta per richiedere l'autenticazione
  },
  { path: '/profile/:userId',
    name: 'Profile',
    component: ProfileView,
    props: true
  },
  { path: '/:pathMatch(.*)*', redirect: '/login' }
]

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes
})

router.beforeEach((to, from, next) => {
  const loggedIn = localStorage.getItem('token')

  if (to.matched.some(record => record.meta.requiresAuth) && !loggedIn) {
    next('/login')
  } else {
    next()
  }
})

export default router
