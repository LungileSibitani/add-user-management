<template>
  <div>
    <h1>Users</h1>

    <!-- Search input -->
    <input
      v-model="searchTerm"
      @input="onSearch"
      placeholder="Search by username or email"
      class="search-input"
    />

    <!-- Users list -->
    <ul>
      <li v-for="user in users" :key="user.id">
        {{ user.username }} - {{ user.email }}
      </li>
    </ul>

    <!-- Load more button -->
    <button @click="loadMore" :disabled="loading || allLoaded">
      {{ loading ? 'Loading...' : allLoaded ? 'All Users Loaded' : 'Load More' }}
    </button>
  </div>
</template>

<script lang="ts">
import { ref, onMounted } from 'vue'

export default {
  setup() {
    const users = ref([])
    const limit = 100
    const offset = ref(0)
    const loading = ref(false)
    const allLoaded = ref(false)
    const searchTerm = ref('')

    const fetchUsers = async (reset = false) => {
      if (loading.value) return
      loading.value = true

      try {
        let url = `http://localhost:8080/users?limit=${limit}&offset=${offset.value}`
        if (searchTerm.value.trim() !== '') {
          url += `&q=${encodeURIComponent(searchTerm.value.trim())}`
        }

        const res = await fetch(url)
        const data = await res.json()

        if (reset) {
          users.value = data
        } else {
          users.value.push(...data)
        }

        if (data.length < limit) {
          allLoaded.value = true
        }
      } catch (err) {
        console.error('Failed to load data:', err)
      } finally {
        loading.value = false
      }
    }

    const loadMore = () => {
      if (allLoaded.value) return
      offset.value += limit
      fetchUsers()
    }

    const onSearch = () => {
      offset.value = 0
      allLoaded.value = false
      fetchUsers(true) // reset users list
    }

    onMounted(() => {
      fetchUsers()
    })

    return { users, loadMore, loading, allLoaded, searchTerm, onSearch }
  },
}
</script>

<style scoped>
.search-input {
  margin-bottom: 10px;
  padding: 5px;
  width: 300px;
}
button {
  margin-top: 10px;
  padding: 5px 10px;
}
</style>
