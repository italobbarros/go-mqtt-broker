document.addEventListener("DOMContentLoaded", function() {
  fetch('http://127.0.0.1:8080/mqtt-tree')
  .then(response => {
      if (!response.ok) {
          throw new Error('Erro ao recuperar a árvore MQTT');
      }
      return response.json();
  })
  .then(data => {
      const treeElement = document.getElementById('tree');
      renderTree(data, treeElement.querySelector('ul'));
  })
  .catch(error => {
      console.error('Erro:', error);
  });

  function renderTree(nodes, parentUl = null) {
    console.log("Renderizando árvore com nodes:", nodes);

    if (!Array.isArray(nodes)) { //significa que é o pai
        console.log("Criando novo parentUl para o nó raiz MQTT");
        const li = document.createElement('li');
        const arrow = document.createElement('span');
        arrow.className = 'arrow';
        if (nodes.children && nodes.children.length > 0){
          li.appendChild(arrow);
        }

        const textNode = document.createTextNode(nodes.name);
        li.appendChild(textNode);

        const ul = document.createElement('ul');
        li.appendChild(ul);

        parentUl.appendChild(li);

        console.log("Elemento anexado ao DOM:", li); // Log para verificar se o elemento está sendo anexado ao DOM

        if (nodes.children && nodes.children.length > 0) {
            renderTree(nodes.children, ul);
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
          if (e.target && e.target.nodeName == "LI") {
              const itemName = e.target.textContent.trim();
              showSidebarDetails(itemName);
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

        const textNode = document.createTextNode(node.name);
        li.appendChild(textNode);

        const ul = document.createElement('ul');
        li.appendChild(ul);

        parentUl.appendChild(li);

        console.log("Elemento anexado ao DOM:", li); // Log para verificar se o elemento está sendo anexado ao DOM

        if (node.children && node.children.length > 0) {
            renderTree(node.children, ul);
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
          if (e.target && e.target.nodeName == "LI") {
              const itemName = e.target.textContent.trim();
              showSidebarDetails(itemName);
          }
        });
    });   
}

function showSidebarDetails(name) {
  const sidebarTitle = document.getElementById('sidebarTitle');
  const sidebarContent = document.getElementById('sidebarContent');

  sidebarTitle.textContent = `Detalhes de ${name}`;
  sidebarContent.textContent = `Informações detalhadas sobre ${name}.`; 
  document.getElementById('sidebar').classList.add('active');
}

});