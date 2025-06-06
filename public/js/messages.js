// Initialisation des conversations
document.addEventListener('DOMContentLoaded', function() {
    // Charger les conversations existantes
    loadConversations();
    
    // Gestion de l'envoi de messages
    document.getElementById('sendMessage')?.addEventListener('click', sendMessage);
});

function loadConversations() {
    // Ici vous devriez charger les conversations depuis votre backend
    // Pour l'exemple, nous créons une conversation par défaut
    const conversationList = document.getElementById('conversationList');
    const defaultConv = createConversationElement('ZdarkBlackShadow', '1');
    conversationList.appendChild(defaultConv);
    
    // Afficher la première conversation par défaut
    if (defaultConv) {
        defaultConv.click();
    }
}

function createConversationElement(name, id) {
    const conv = document.createElement('div');
    conv.className = 'conversation';
    conv.setAttribute('data-conversation-id', id);
    conv.innerHTML = `
        <div class="conversation-avatar">${name.substring(0, 2)}</div>
        <div class="conversation-name">${name}</div>
    `;
    
    conv.addEventListener('click', function() {
        // Activer la conversation
        document.querySelectorAll('.conversation').forEach(c => c.classList.remove('active'));
        this.classList.add('active');
        
        // Afficher la zone de chat
        document.getElementById('chatArea').style.display = 'flex';
        
        // Mettre à jour l'en-tête
        document.getElementById('chatName').textContent = name;
        document.getElementById('chatAvatar').textContent = name.substring(0, 2);
        
        // Charger les messages (à remplacer par un appel à votre backend)
        loadMessages(id);
    });
    
    return conv;
}

function loadMessages(conversationId) {
    // Ici vous devriez charger les messages depuis votre backend
    const chatMessages = document.getElementById('chatMessages');
    chatMessages.innerHTML = `
        <div class="message">
            <div class="message-avatar">${document.getElementById('chatAvatar').textContent}</div>
            <div class="message-content">
                <div class="message-author">${document.getElementById('chatName').textContent}</div>
                <div class="message-text">Ceci est un message de démo.</div>
            </div>
        </div>
    `;
}

function sendMessage() {
    const messageInput = document.getElementById('messageInput');
    const message = messageInput.value.trim();
    
    if (message) {
        const chatMessages = document.getElementById('chatMessages');
        const activeConversation = document.querySelector('.conversation.active');
        
        if (activeConversation) {
            // Ajouter le nouveau message
            const messageElement = document.createElement('div');
            messageElement.className = 'message';
            messageElement.innerHTML = `
                <div class="message-avatar">M</div>
                <div class="message-content">
                    <div class="message-author">Moi</div>
                    <div class="message-text">${message}</div>
                </div>
            `;
            
            chatMessages.appendChild(messageElement);
            messageInput.value = '';
            
            // Faire défiler vers le bas
            chatMessages.scrollTop = chatMessages.scrollHeight;
            
            // Ici vous devriez envoyer le message à votre backend
            console.log('Message envoyé:', message);
        }
    }
}

function acceptMessageRequest() {
    const popup = document.getElementById('messagePopup');
    popup.style.display = 'none';
    
    // Créer une nouvelle conversation
    const conversationList = document.getElementById('conversationList');
    const newConv = createConversationElement('Nouvel Ami', Date.now());
    conversationList.appendChild(newConv);
    
    newConv.click();
}

function rejectMessage() {
    document.getElementById('messagePopup').style.display = 'none';
}

fetch('/public/txt/emoji.txt')
    .then(response => response.text())
    .then(data => {
        const emojiList = document.getElementById('emojiList');
        const emojis = data.split(/\s+/).filter(e => e.trim().length > 0);
        emojis.forEach(emoji => {
            const span = document.createElement('span');
            span.textContent = emoji;
            span.className = 'emoji-item';
            span.style.cursor = 'pointer';
            span.onclick = function() {
                const input = document.getElementById('messageInput');
                input.value += emoji;
                input.focus();
            };
            emojiList.appendChild(span);
        });
    })
    .catch(err => {
        console.error("Erreur de chargement des emojis:", err);
    });