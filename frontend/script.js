document.addEventListener("DOMContentLoaded", function() {
  fetch('http://127.0.0.1:8080/mqtt-tree')
  .then(response => {
      if (!response.ok) {
          throw new Error('Erro ao recuperar a árvore MQTT');
      }
      return response.json();
  })
  .then(data => renderTree(data))
  .catch(error => {
      console.error('Erro:', error);
  });

  function renderTree(node, parentElement) {
    const ul = document.createElement('ul');

    node.forEach(item => {
        const li = document.createElement('li');
        const arrow = document.createElement('span');
        arrow.className = 'arrow';
        arrow.textContent = '▶'; // ou qualquer ícone de seta que você esteja usando
        li.appendChild(arrow);

        const textNode = document.createTextNode(item.name);
        li.appendChild(textNode);

        if (item.children && item.children.length > 0) {
            renderTree(item.children, li); // Renderiza os filhos recursivamente
        }

        ul.appendChild(li);
    });

    parentElement.appendChild(ul);
}

  function createNode(nodeData) {
      const li = document.createElement('li');
      const span = document.createElement('span');
      span.textContent = nodeData.Name;

      if (nodeData.Children && nodeData.Children.length > 0) {
          span.addEventListener('click', function() {
              toggleChildren(li);
          });
          li.appendChild(span);
          const ul = document.createElement('ul');
          li.appendChild(ul);
          nodeData.Children.forEach(child => {
              ul.appendChild(createNode(child));
          });
      } else {
          li.textContent = nodeData.Name;
      }

      return li;
  }

  
  function toggleChildren(node) {
      const ul = node.querySelector('ul');
      if (ul) {
          ul.classList.toggle('collapsed');
      }
  }
});

// script.js

document.addEventListener("DOMContentLoaded", function() {
  const folders = document.querySelectorAll('.folder');

  folders.forEach(folder => {
      const arrow = folder.querySelector('.arrow');
      const childUl = folder.querySelector('ul');
      
      arrow.addEventListener('click', function(e) {
          e.stopPropagation(); // Impede que o evento seja propagado para os elementos pai
          
          if (childUl) {
              childUl.classList.toggle('hidden'); // Toggle para esconder ou mostrar os filhos
          }
          
          // Trocar a classe do folder para alterar a seta
          folder.classList.toggle('collapsed');
          folder.classList.toggle('expanded');
      });
  });
});
