
document.getElementById('loginForm').addEventListener('submit', async (e) => {

    e.preventDefault();

    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;

    try {
        const response = await fetch('http://localhost:8000/user/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ username, password })
        });

        if (response.ok) {
            const data = await response.json();
            const token = data.token;

            localStorage.setItem('jwtToken', token);
            localStorage.setItem('user',username)
            window.location.href = '../devices/devices.html';
        } else {
            const error = await response.json();
            document.getElementById('message').textContent = 'Invalid username/password';
        }
    } catch (error) {
        document.getElementById('message').textContent = 'Invalid username/password';
    }
});
