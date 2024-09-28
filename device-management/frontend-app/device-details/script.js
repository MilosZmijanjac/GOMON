document.addEventListener('DOMContentLoaded', async () => {
    const device =localStorage.getItem('selectedDevice');
    const token = localStorage.getItem('jwtToken');
    if (!token) {
        window.location.href = '../login/index.html';
        return;
    }
    if (!device) {
        window.location.href = '../devices/devices.html';
        return;
    }
    selectedDevice=JSON.parse(device);
    // Fetch device details
    const response = await fetch(`http://localhost:8000/device/telemetry`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${localStorage.getItem('jwtToken')}`
        },
        body:JSON.stringify({DeviceID:selectedDevice.DeviceID})
    });
    const data = await response.json();

    const cpuUsage = data.cpuUsages;
    const temperatures = data.temperatures;
    const timestamps =temperatures? data.timestamps.map(entry => new Date(entry).toLocaleTimeString()):[];

    new Chart(document.getElementById('cpuChart'), {
        type: 'line',
        data: {
            labels: timestamps,
            datasets: [{
                label: 'CPU Usage (%)',
                data: cpuUsage,
                borderColor: 'rgba(75, 192, 192, 1)',
                backgroundColor: 'rgba(75, 192, 192, 0.2)',
                fill: true
            }]
        },
        options: {
            responsive: true,
            scales: {
                x: {
                    beginAtZero: true
                },
                y: {
                    beginAtZero: true,
                    max: 100
                }
            }
        }
    });

    new Chart(document.getElementById('temperatureChart'), {
        type: 'line',
        data: {
            labels: timestamps,
            datasets: [{
                label: 'Temperature (°C)',
                data: temperatures,
                borderColor: 'rgba(255, 99, 132, 1)',
                backgroundColor: 'rgba(255, 99, 132, 0.2)',
                fill: true
            }]
        },
        options: {
            responsive: true,
            scales: {
                x: {
                    beginAtZero: true
                },
                y: {
                    beginAtZero: true
                }
            }
        }
    });

    document.getElementById('restartButton').addEventListener('click', async() => {
        const response = await fetch(`http://localhost:8000/device/command`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${localStorage.getItem('jwtToken')}`
            },
            body:JSON.stringify({DeviceID:selectedDevice.DeviceID,Command:1})
        });
    });

    document.getElementById('shutdownButton').addEventListener('click', async() => {
        const response = await fetch(`http://localhost:8000/device/command`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${localStorage.getItem('jwtToken')}`
            },
            body:JSON.stringify({DeviceID:selectedDevice.DeviceID,Command:2})
        });
    });

    async function loadNotifications() {
        
        const response = await fetch(`http://localhost:8000/device/notification`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${localStorage.getItem('jwtToken')}`
            },
            body:JSON.stringify({DeviceID:selectedDevice.DeviceID,})
        });
        var dataJson=await response.json();
        const notificationContent = document.querySelector('.notification-content');
        notificationContent.innerHTML = dataJson.Notifications.map(notif => `<p>${new Date(notif.timestamp).toLocaleTimeString()} | ${notif.code}|</p>`).join('');
    }

    await loadNotifications();
});

