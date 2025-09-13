<template>
  <div>
    <h1>Users</h1>
    <ul>
      <li v-for="user in users" :key="user.id">
        {{ user.username }} - {{ user.email }}
      </li>
    </ul>
  </div>
</template>

<script lang="ts">
import { ref, onMounted } from 'vue'

export default {
  setup() {
    const users = ref([])

    onMounted(async () => {
      try {
        const res = await fetch('http://localhost:8080/users')
        users.value = await res.json()
      } catch (err) {
        console.error('Failed to load data:', err)
      }
    })

    return { users }
  }
}
</script>
