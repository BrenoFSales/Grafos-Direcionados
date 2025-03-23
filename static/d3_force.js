const width = window.innerWidth;
const height = window.innerHeight;

const svg = d3.select("#grafo-exibicao").append("svg")
    .attr("width", width)
    .attr("height", height);

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

// Declaração dos nós do dígrafo	
const nodes = [
    { id: "A" }, { id: "B" }, { id: "C" }, { id: "D" },
    { id: "E" }
];

// Declaração das conexões direcionadas de cada nó
const links = [
    { source: "A", target: "B" },
    { source: "A", target: "E" },
    { source: "A", target: "C" },
    { source: "B", target: "D" },
    { source: "C", target: "B" },
    { source: "D", target: "A" },
    { source: "D", target: "C" },
    { source: "D", target: "E" },
    { source: "E", target: "C" }
];

// Criação da Simulação (Forças)
const simulation = d3.forceSimulation(nodes)
    .force("link", d3.forceLink(links).id(d => d.id).distance(300))
    .force("charge", d3.forceManyBody().strength(-400))
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