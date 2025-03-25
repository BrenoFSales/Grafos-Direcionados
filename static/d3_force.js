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

	let input = document.querySelector('input#nameNode');
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
	await atualizarMatrizAdjacencia();

	input.value = '';
}

async function deletarNode() {
	// deleta o nó no grafo, e então sincroniza com o back.
	let input = document.querySelector('input#nameNode');
	let exemplo = document.querySelector('select#preset');
	input = input;
	input.value = input.value;

	let deletar = { id: input.value };

	let resposta = await fetch(`/node/${exemplo.value}`, {method: 'DELETE', body: JSON.stringify(deletar)});
	if (!resposta.ok) {
		throw resposta.ok;
	}

	links = links.filter(link => !(link.source.id === input.value || link.target.id === input.value));
	nodes = nodes.filter(node => node.id !== deletar.id);

	renderizar(nodes, links);

	atualizarListaDeNodes();

	input.value = '';

	// OBS! Se adicionar um novo Nó e removê-lo só funciona depois de recarregar o cache,
	// e caso o Nó tenha arestas ligadas a ele, ao removê-lo a aplicação crasha
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

async function deletarAresta() {
	let de = document.querySelector('#select-de');
	let para = document.querySelector('#select-para');
	let exemplo = document.querySelector('select#preset');

	var link = { source: de.value, target: para.value };

	let resposta = await fetch(`/link/${exemplo.value}`, { method: 'DELETE', body: JSON.stringify(link) });
	if (!resposta.ok) {
		throw resposta.ok;
	}

	let index = links.findIndex(({ source, target }) => source.id == link.source && target.id == link.target);
	if (index < 0) { throw index; } links = links.filter((_, i) => i != index);

	renderizar(nodes, links);
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
	await atualizarMatrizAdjacencia();
}

// isso faz questão de manter salvo o exemplo selecionado para que o usuário
// veja o mesmo exemplo toda vez que recarregue a página.
async function trocarExemplo(evento) {

	const url = new URL(window.location.href);

	//
	let exemploSelecionado = document.querySelector('select#preset').value;
	let parametroExemploSalvo = url.searchParams.get('exemplo');

	let casoParametroExemploPresente = parametroExemploSalvo !== null;
	let casoUsuarioRecarregouPagina = evento === undefined;
	// evento será um objeto defindo caso essa função tenha sido chamada pelo onchange do elemento select.
	// nesse caso, o usuário selecionou um novo exemplo e queremos mudar para o novo exemplo.
	// evento é undefined quando essa função é chamada uma única vez no primeiro load da página.
	// nesse caso, queromos selecionar o exemplo salvo na url.
	if (casoUsuarioRecarregouPagina && casoParametroExemploPresente) {
		exemploSelecionado = parametroExemploSalvo;
		// alterar select para que também reflita a escolha.
		document.querySelector('select#preset').value = parametroExemploSalvo;
	}

	let resposta = await fetch(`/node/${exemploSelecionado}`, { method: 'GET' });
	if (!resposta.ok) {
		throw resposta.ok;
	}
	let { nodes: nodes_, links: links_ } = await resposta.json();
	nodes = nodes_;
	links = links_;

	renderizar(nodes, links);
	atualizarListaDeNodes();
	atualizarMatrizAdjacencia();

	// atualiza a url para que toda vez que o usuário recarregar a página, o mesmo exemplo será exibido.
	const parametrosNovos = new URLSearchParams({ exemplo: exemploSelecionado }).toString();

	console.log(parametrosNovos);
	window.history.replaceState(null, "", `${url.pathname}?${parametrosNovos}`)

}

trocarExemplo();

async function atualizarMatrizAdjacencia() {
	let exibirRotulos = document.querySelector('#exibir-rotulos');
	let exemplo = document.querySelector('select#preset');
	let resposta = await fetch(`/matriz/${exemplo.value}?rotulo=${exibirRotulos.checked}`);
	if (!resposta.ok) {
		throw resposta.ok;
	}
	let matriz = await resposta.text();
	let latex = `\\begin{bmatrix}\n${matriz}\n\\end{bmatrix}`;
	katex.render(latex, document.querySelector('#matriz'), { throwOnError: true, });
}

async function toggleMatrizAdjacencia() {
	let matriz = document.querySelector('#matriz');
	let butao = document.querySelector('#matriz-toggle');
	matriz.classList.toggle('hidden')
	if (matriz.classList.contains('hidden')) {
		butao.textContent = 'Mostrar';
	} else {
		butao.textContent = 'Esconder';
		await atualizarMatrizAdjacencia();
	}
}
