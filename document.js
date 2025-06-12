document.addEventListener('DOMContentLoaded', function() {
    loadConversations();
    document.getElementById('sendMessage')?.addEventListener('click', sendMessage);
});

function loadConversations() {
    const conversationList = document.getElementById('conversationList');
    const defaultConv = createConversationElement('ZdarkBlackShadow', '1');
    conversationList.appendChild(defaultConv);
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
        document.querySelectorAll('.conversation').forEach(c => c.classList.remove('active'));
        this.classList.add('active');
        document.getElementById('chatArea').style.display = 'flex';
        document.getElementById('chatName').textContent = name;
        document.getElementById('chatAvatar').textContent = name.substring(0, 2);
        loadMessages(id);
    });
    
    return conv;
}

function loadMessages(conversationId) {
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
            
            chatMessages.scrollTop = chatMessages.scrollHeight;
            console.log('Message envoyé:', message);
        }
    }
}

function acceptMessageRequest() {
    const popup = document.getElementById('messagePopup');
    popup.style.display = 'none';
    const conversationList = document.getElementById('conversationList');
    const newConv = createConversationElement('Nouvel Ami', Date.now());
    conversationList.appendChild(newConv);
    
    newConv.click();
}

function rejectMessage() {
    document.getElementById('messagePopup').style.display = 'none';
}

function loadEmojisDynamic() {
    const emojiList = document.getElementById('emojiList');
    emojiList.innerHTML = '';
    fetch('/public/txt/emoji.txt')
        .then(response => response.text())
        .then(data => {
            const emojis = data.split(/\s+/).filter(e => e.trim().length > 0);
            emojis.forEach(emoji => {
                const span = document.createElement('span');
                span.textContent = emoji;
                span.classList.add('emoji-item');
                span.style.cursor = 'pointer';
                span.addEventListener('click', () => {
                    const input = document.getElementById('messageInput');
                    input.value += emoji;
                    input.focus();
                });
                emojiList.appendChild(span);
            });
        })
        .catch(err => {
            console.error("Erreur de chargement des emojis:", err);
            emojiList.textContent = "Erreur de chargement";
        });
}

document.addEventListener("DOMContentLoaded", function() {
    const emojiPanel = document.getElementById('emojiPanel');
    if(emojiPanel) {
        loadEmojisDynamic();
    }
});