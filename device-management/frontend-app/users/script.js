document.addEventListener('DOMContentLoaded', async () => {
    const logoutButton = document.getElementById('logout');
    const addNewUser=document.getElementById('addNewUser');
    if (logoutButton) {
        logoutButton.addEventListener('click', function() {
            localStorage.removeItem('jwtToken');
            localStorage.removeItem('user');
            localStorage.removeItem('selectedDevice');
            localStorage.removeItem('selectedUser');
            window.location.href = '../login/index.html';
        });
    }
    if (addNewUser) {
        addNewUser.addEventListener('click', function() {
            window.location.href = '../add-user/add-user.html'; 
        });
    }
})
document.addEventListener('DOMContentLoaded', async () => {
    const token = localStorage.getItem('jwtToken');
    if (!token) {
        window.location.href = '../login/index.html';
        return;
    }
    const tableBody = document.getElementById('usersTable').getElementsByTagName('tbody')[0];
    
        const response = await fetch('http://localhost:8000/user/list', {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });

        if (response.ok) {
            const users = await response.json();
            users.users.forEach(user => {
                const row = tableBody.insertRow();

                row.insertCell(0).textContent = user.Username;
                row.insertCell(1).textContent = user.Role==0?'Admin':'User';
                row.insertCell(2).textContent = user.IsActive?'Active':'Inactive';
                  row.addEventListener('click', function() {
                     localStorage.setItem('selectedUser',JSON.stringify(user))
                     window.location.href = "../edit-user/edit-user.html";
                });
            });
        } else {
            console.error('Error fetching users');
            window.location.href = '../login/index.html';
        }

});