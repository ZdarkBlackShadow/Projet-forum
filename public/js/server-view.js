document.addEventListener('DOMContentLoaded', () => {
    const params = new URLSearchParams(window.location.search);
    const serverName = params.get('name') || "Serveur par défaut";
    
    const serverNameDisplay = document.getElementById('serverNameDisplay');
    if (serverNameDisplay) {
        serverNameDisplay.textContent = serverName;
    }
    document.title = serverName; // Set page title

    const channels = document.querySelectorAll('.channel-item');
    const currentChannelNameDisplay = document.getElementById('currentChannelName');
    const messageInputServer = document.getElementById('messageInputServer');

    channels.forEach(channel => {
        channel.addEventListener('click', () => {
            channels.forEach(c => c.classList.remove('active'));
            channel.classList.add('active');
            const channelName = channel.textContent.replace('#', '').trim();
            if (currentChannelNameDisplay) {
                currentChannelNameDisplay.textContent = channelName;
            }
            if (messageInputServer) {
                messageInputServer.placeholder = `Envoyer un message dans #${channelName}`;
            }
            // Add logic here to load messages for the selected channel
            console.log(`Switched to channel: ${channelName}`);
        });
    });

    const serverSettingsButton = document.getElementById('serverSettingsButton');
    if (serverSettingsButton) {
        serverSettingsButton.addEventListener('click', () => {
            // Placeholder for server settings functionality
            // This would typically open a modal or navigate to a settings page
            alert(`Paramètres pour le serveur : ${serverName}\n(Fonctionnalité de modification à implémenter)`);
            // Example: prompt for new server name
            // const newName = prompt("Entrez le nouveau nom du serveur:", serverName);
            // if (newName && newName.trim() !== "") {
            //     serverNameDisplay.textContent = newName.trim();
            //     document.title = newName.trim();
            //     // Here, an API call would be made to save the new server name
            // }
        });
    }
    
    // Simulate sending a message
    const messageInput = document.getElementById('messageInputServer');
    if (messageInput) {
        messageInput.addEventListener('keypress', function(e) {
            if (e.key === 'Enter' && messageInput.value.trim() !== '') {
                const chatMessagesServer = document.getElementById('chatMessagesServer');
                const newMessage = document.createElement('div');
                newMessage.classList.add('message-item-server');
                newMessage.innerHTML = `
                    <img src="/public/img/default_image.png" alt="Avatar" class="message-avatar-server">
                    <div class="message-content-server">
                        <span class="message-author-server">VotrePseudo <span class="message-timestamp">Maintenant</span></span>
                        <p class="message-text-server">${messageInput.value.trim()}</p>
                    </div>
                `;
                chatMessagesServer.appendChild(newMessage);
                chatMessagesServer.scrollTop = chatMessagesServer.scrollHeight; // Scroll to bottom
                messageInput.value = '';
                e.preventDefault();
            }
        });
    }
});
