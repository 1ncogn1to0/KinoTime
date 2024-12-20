document.addEventListener('DOMContentLoaded', async () => {
    const users = await apiClient.getAllUsers();
    console.log(users);
});
