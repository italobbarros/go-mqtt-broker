document.addEventListener("DOMContentLoaded", function() {
  fetch('http://192.168.0.109:8080/mqtt-tree')
  .then(response => {
      if (!response.ok) {
          throw new Error('Erro ao recuperar a árvore MQTT');
      }
      return response.json();
  })
  .then(data => {
      const treeElement = document.getElementById('tree');
      renderTree(data,data, treeElement.querySelector('ul'));
  })
  .catch(error => {
      console.error('Erro:', error);
  });

  function renderTree(nodes,dataObject, parentUl = null) {
    console.log("Renderizando árvore com nodes:", nodes);
    if (!Array.isArray(nodes)) { //significa que é o pai
      console.log("Criando novo parentUl para o nó raiz MQTT");
      const li = document.createElement('li');
      const arrow = document.createElement('span');
      arrow.className = 'arrow';
      if (nodes.children && nodes.children.length > 0){
        li.appendChild(arrow);
      }

      const textNode = document.createTextNode("/"+nodes.name);
      li.appendChild(textNode);

      const ul = document.createElement('ul');
      li.appendChild(ul);

      parentUl.appendChild(li);

      console.log("Elemento anexado ao DOM:", li); // Log para verificar se o elemento está sendo anexado ao DOM

      if (nodes.children && nodes.children.length > 0) {
          renderTree(nodes.children,dataObject, ul);
      }
      
      arrow.addEventListener('click', function() {
        console.log("Clique na seta!"); // Log para verificar se o evento de clique está sendo acionado
        if (ul.classList.toggle('hidden')) {
            arrow.classList.toggle('collapsed', true);
            arrow.classList.toggle('expanded', false);
        } else {
            arrow.classList.toggle('expanded', true);
            arrow.classList.toggle('collapsed', false);
            console.log("expanded");
        }
      });
      li.addEventListener('click', function(e) {
        if (e.target && e.target.nodeName === "LI") {
            const clickedItem = getCurrentClickText(e.target);
            handleListItemClick(dataObject,clickedItem);
        }
      });
      
      return parentUl;
    }

    nodes.forEach(node => {
        console.log("Renderizando node:", node); // Log para verificar o node sendo processado

        const li = document.createElement('li');
        const arrow = document.createElement('span');
        arrow.className = 'arrow';
        if (node.children && node.children.length > 0){
          li.appendChild(arrow);
        }

        const textNode = document.createTextNode("/"+node.name);
        li.appendChild(textNode);

        const ul = document.createElement('ul');
        li.appendChild(ul);

        parentUl.appendChild(li);

        console.log("Elemento anexado ao DOM:", li); // Log para verificar se o elemento está sendo anexado ao DOM

        if (node.children && node.children.length > 0) {
            renderTree(node.children,dataObject, ul);
        }
        
        arrow.addEventListener('click', function() {
          console.log("Clique na seta!"); // Log para verificar se o evento de clique está sendo acionado
          if (ul.classList.toggle('hidden')) {
              arrow.classList.toggle('collapsed', true);
              arrow.classList.toggle('expanded', false);
          } else {
              arrow.classList.toggle('expanded', true);
              arrow.classList.toggle('collapsed', false);
          }
        });
    });   
}

function handleListItemClick(dataObject,name) {
  const topic = findTopicByName(dataObject, name); // Substitua yourDataObject pelo objeto de dados retornado do endpoint
  if (topic) {
    console.log("topic: "+topic)
    showSidebarDetails(topic); // Chama a função com o tópico correspondente
  } else {
    console.log("Tópico não encontrado para o nome:", name);
  }
}

function findTopicByName(data, name) {
  if (data.name === name) {
      return data.topic;
  }
  if (data.children) {
      for (let child of data.children) {
          const result = findTopicByName(child, name);
          if (result) return result;
      }
  }
  return null; // Se não encontrou, retorne null
}

function getCurrentClickText(element) {
  // Pega o texto do elemento e remove qualquer conteúdo após o primeiro <ul>
  const textContent = element.textContent.trim();
    // Remove qualquer conteúdo após o primeiro '/'
  const index = textContent.split('/');
  console.log(index)
  return index[0]=="" ? index[1] : index[0];
  
}

function processTopic(topic) {
  // Verifica se a string contém '#'
  if (topic.includes('#')) {
    // Substitui todos os '#' por '%23'
    return topic.replace(/#/g, '%23');
  }
  // Retorna a string original se não houver '#'
  return topic;
}

function showSidebarDetails(name) {
  console.log("showSidebarDetails")
  const sidebarTitle = document.getElementById('sidebarTitle');
  const sidebarContent = document.getElementById('sidebarContent');
  sidebarContent.innerHTML = '';

  sidebarTitle.textContent = `Detalhes de ${name}`;
  //sidebarContent.textContent = `Informações detalhadas sobre ${name}.`; 
  
  fetch(`http://192.168.0.109:8080/topic-info?topic=${processTopic(name)}`)
  .then(response => {
      if (!response.ok) {
          throw new Error('Erro ao recuperar as informações do tópico');
      }
      return response.json();
  })
  .then(data => {
    console.log("aqui")
    renderTopicInfo(data);
  })
  .catch(error => {
      console.error('Erro:', error);
  });

  function renderTopicInfo(topicInfo) {
    const sidebar = document.getElementById('sidebarContent');
    sidebarContent.innerHTML = '';
    // Criar elementos para mostrar as informações do tópico
    const topicName = document.createElement('h2');
    topicName.textContent = `Tópico: ${topicInfo.topic}`;

    const messageCount = document.createElement('p');
    messageCount.textContent = `Número de Mensagens: ${topicInfo.messageCount}`;

    const payload = document.createElement('p');
    payload.textContent = `Payload: ${topicInfo.topicCfg.payload}`;

    const qos = document.createElement('p');
    qos.textContent = `Qos: ${topicInfo.topicCfg.qos}`;

    const retained = document.createElement('p');
    retained.textContent = `Retained: ${topicInfo.topicCfg.retained}`;

    const subscribers = document.createElement('p');
    const count = topicInfo.topicCfg.subscribers ? topicInfo.topicCfg.subscribers.length : 0;
    subscribers.textContent = `Número de Inscritos: ${count}`;
    // Adicionar elementos à sidebar
    sidebar.appendChild(topicName);
    sidebar.appendChild(messageCount);
    sidebar.appendChild(payload);
    sidebar.appendChild(qos);
    sidebar.appendChild(retained);
    sidebar.appendChild(subscribers);
  }
  document.getElementById('sidebar').classList.add('active');
}

});

