import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import LoginView from '../views/LoginView.vue'
import ProfileView from '@/views/ProfileView.vue'
import NonLoggedProfileView from '@/views/NonLoggedProfileView.vue'
import UploadPhoto from '@/views/UploadPhotoView.vue'; 

const routes = [
  { path: '/login', component: LoginView },
  { 
    path: '/', 
    component: HomeView,
    meta: { requiresAuth: true } // Aggiungi meta per richiedere l'autenticazione
  },
  { path: '/user/:userId/profile',
    name: 'Profile',
    component: ProfileView,
    props: true
  },
  { path: '/:pathMatch(.*)*', redirect: '/login' },
  { 
    path: '/user/:userId/nonlogged-profile', 
    name: 'NonLoggedProfile', 
    component: NonLoggedProfileView,
    props: true 
  },
  {
    path: '/search',
    name: 'SearchUser',
    component: () => import('../views/SearchView.vue')  // Assumi che esista una view dedicata alla ricerca
  },
  {
    path: '/upload',
    name: 'UploadPhoto',
    component: UploadPhoto
  }
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
