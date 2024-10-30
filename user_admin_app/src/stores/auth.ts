import { defineStore } from 'pinia'
import { computed, ref } from 'vue'
import { useRouter } from 'vue-router'

export const useAuthStore = defineStore('auth', () => {
  const router = useRouter()
  const session = ref('')

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
      setSession(cookie.value, cookie.expires)
    }
  }

  function setSession(sessionStr: string, expires: number) {
    const now = new Date()
    const diff = expires - now.getTime()

    // If we want to ask for a session refresh
    // setTimeout(
    //   () => {
    //     confirm('Session almost finished')
    //   },
    //   diff - 30 * 1000
    // )

    session.value = sessionStr
    setTimeout(() => {
      clearSession()
    }, diff)
    router.push('/about')
  }

  function clearSession() {
    session.value = ''
    router.push('/')
  }

  return { isLoggedIn, init, clearSession, setSession }
})
