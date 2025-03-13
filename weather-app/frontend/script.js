// function getWeather() {
//     const city = document.getElementById('city').value;
//     if (!city) {
//         alert('Please enter a city name');
//         return;
//     }

//     fetch(`http://localhost:8080/weather?city=${city}`)
//         .then(response => response.json())
//         .then(data => {
//             const weatherResult = document.getElementById('weather-result');
//             weatherResult.innerHTML = `
//                 <h2>${data.location.name}, ${data.location.country}</h2>
//                 <p>Temperature: ${data.current.temp_c}°C</p>
//                 <p>Condition: ${data.current.condition.text}</p>
//             `;
//         })
//         .catch(error => {
//             console.error('Error:', error);
//         });
// }

function getWeather() {
    const city = document.getElementById('city').value;
    if (!city) {
        alert('Please enter a city name');
        return;
    }

    // URL backend (pastikan ini sesuai dengan yang berjalan)
    const apiUrl = `http://localhost:8080/weather?city=${city}`;

    fetch(apiUrl)
        .then(response => {
            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            const weatherResult = document.getElementById('weather-result');
            if (data && data.location && data.current) {
                weatherResult.innerHTML = `
                    <h2>${data.location.name}, ${data.location.country}</h2>
                    <p>Temperature: ${data.current.temp_c}°C</p>
                    <p>Condition: ${data.current.condition.text}</p>
                `;
            } else {
                weatherResult.innerHTML = `<p>No weather data found for ${city}.</p>`;
            }
        })
        .catch(error => {
            console.error('Error:', error);
            const weatherResult = document.getElementById('weather-result');
            weatherResult.innerHTML = `<p>Failed to fetch weather data. Please try again later.</p>`;
        });
}