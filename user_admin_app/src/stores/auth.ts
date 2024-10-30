import { defineStore } from 'pinia'
import { computed, ref } from 'vue'
import { useRouter } from 'vue-router'

export const useAuthStore = defineStore('auth', () => {
  const session = ref('')
  const router = useRouter()

  const isLoggedIn = computed(() => !!session.value)
  //session.value -> string = "12345..." -> truthy
  // -> con el primer bang pasa a false
  // -> con el segundo bang pasa a true
  //session.value -> string = "" -> falsey
  // -> con el primer bang pasa a true
  // -> con el segundo bang pasa a false

  async function init() {
    const cookie = await (window as any).cookieStore.get('session')
    if (cookie) {
      setSession(cookie.value, 'about')
    }
  }

  function setSession(sessionStr: string, goTo: string) {
    session.value = sessionStr
    if (goTo) router.push({ name: goTo })
  }

  function clearSession() {
    session.value = ''
  }

  return { init, isLoggedIn, clearSession, setSession }
})
