var width = window.innerWidth;
var height = window.innerHeight;

window.onresize = () => {
	width = window.innerWidth;
	height = window.innerHeight;
	let svg = document.querySelector('svg');
	svg.setAttribute("width", width);
	svg.setAttribute("height", height);
}

// Declaração dos nós do dígrafo
var nodes = [
	{ id: "A" }, { id: "B" }, { id: "C" }, { id: "D" },
	{ id: "E" }
];

// Declaração das conexões direcionadas de cada nó
var links = [
	{ source: "A", target: "E" },
	{ source: "B", target: "D" },
	{ source: "C", target: "B" },
	{ source: "D", target: "A" },
	{ source: "D", target: "C" },
	{ source: "D", target: "E" },
	{ source: "E", target: "C" },
	{ source: "E", target: "C" },
];
// D3 usa a referência desses dois objetos para salvar informações essenciais dos nós como posição e etc.

function renderizar(nodes, links) {

	// Limpa qualquer coisa que tenha sido feito previamente para
	// recriar o grafo.
	document.querySelector("#grafo-exibicao").innerHTML = '';

	const svg = d3.select("#grafo-exibicao").append("svg");
	// .attr("width", width)
	// .attr("height", height);

	// Define as setas das arestas
	svg.append("defs").append("marker")
		.attr("id", "arrow")
		.attr("viewBox", "0 -5 10 10")
		.attr("refX", 18)
		.attr("refY", 0)
		.attr("markerWidth", 17)
		.attr("markerHeight", 17)
		.attr("orient", "auto")
		.append("path")
		.attr("d", "M0,-5L10,0L0,5")
		.attr("fill", "#9dbaea");

	// Criação da Simulação (Forças)
	const simulation = d3.forceSimulation(nodes)
		.force("link", d3.forceLink(links).id(d => d.id).distance(300))
		.force("charge", d3.forceManyBody().strength(-60))
		.force("center", d3.forceCenter(width / 2, height / 2));

	const link = svg.selectAll(".link")
		.data(links)
		.enter().append("line")
		.attr("class", "link");

	const node = svg.selectAll(".node")
		.data(nodes)
		.enter().append("g")
		.attr("class", "node");

	// Tamanho dos nós
	node.append("circle")
		.attr("r", 30);

	// Tamanho do texto dentro do nó
	node.append("text")
		.text(d => d.id)
		.attr("dy", 5);

	simulation.on("tick", () => {
		link
			.attr("x1", d => d.source.x)
			.attr("y1", d => d.source.y)
			.attr("x2", d => d.target.x)
			.attr("y2", d => d.target.y);

		node
			.attr("transform", d => `translate(${d.x},${d.y})`);
	});

	// Adiciona a física de arraso e repulsão dos nós
	const drag = d3.drag()
		.on("start", (event, d) => {
			if (!event.active) simulation.alphaTarget(0.3).restart();
			d.fx = d.x;
			d.fy = d.y;
		})
		.on("drag", (event, d) => {
			d.fx = event.x;
			d.fy = event.y;
		})
		.on("end", (event, d) => {
			if (!event.active) simulation.alphaTarget(0);
			d.fx = null;
			d.fy = null;
		});

	// Inicializa a mecânica de física
	node.call(drag);
}

async function adicionarNode() {

	// adiciona um nó ao grafo, porém antes sincroniza com o back todos os nós que temos aqui.

	let input = document.querySelector('input#adicionar');
	let exemplo = document.querySelector('select#preset');
	input = input;
	input.value = input.value;

	let novo = { id: input.value, x: width / 2, y: height / 2 };

	let resposta = await fetch(`/node/${exemplo.value}`, { method: 'POST', body: JSON.stringify(novo) });
	if (!resposta.ok) {
		throw resposta.ok;
	}

	nodes.push(novo);

	renderizar(nodes, links);

	atualizarListaDeNodes();

	input.value = '';
}

function atualizarListaDeNodes() {
	let de = document.querySelector('#select-de');
	let para = document.querySelector('#select-para');
	de.innerHTML = '';
	para.innerHTML = '';

	for (let i = 0; i < nodes.length; i++) {
		de.innerHTML += `<option value="${nodes[i].id}">${nodes[i].id}</option>`;
		para.innerHTML += `<option value="${nodes[i].id}">${nodes[i].id}</option>`;
	}
}

async function conectarNodes() {
	let de = document.querySelector('#select-de');
	let para = document.querySelector('#select-para');
	let exemplo = document.querySelector('select#preset');

	var link = { source: de.value, target: para.value };


	let resposta = await fetch(`/link/${exemplo.value}`, { method: 'POST', body: JSON.stringify(link) });
	if (!resposta.ok) {
		throw resposta.ok;
	}

	links.push(link);

	renderizar(nodes, links);
}

async function trocarExemplo() {
	let exemplo = document.querySelector('select#preset');

	let resposta = await fetch(`/node/${exemplo.value}`, { method: 'GET' });
	if (!resposta.ok) {
		throw resposta.ok;
	}
	let { nodes: nodes_, links: links_ } = await resposta.json();
	nodes = nodes_;
	links = links_;
	renderizar(nodes, links);
	atualizarListaDeNodes();
}

trocarExemplo();
