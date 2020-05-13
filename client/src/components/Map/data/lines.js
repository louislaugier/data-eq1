import line1 from '../../../assets/lines/line1.png'
import line2 from '../../../assets/lines/line2.png'
import line3 from '../../../assets/lines/line3.png'
import line3bis from '../../../assets/lines/line3bis.png'
import line4 from '../../../assets/lines/line4.jpeg'
import line5 from '../../../assets/lines/line5.png'
import line6 from '../../../assets/lines/line6.png'
import line7 from '../../../assets/lines/line7.png'
import line7bis from '../../../assets/lines/line7bis.png'
import line8 from '../../../assets/lines/line8.png'
import line9 from '../../../assets/lines/line9.png'
import line10 from '../../../assets/lines/line10.png'
import line11 from '../../../assets/lines/line11.png'
import line12 from '../../../assets/lines/line12.png'
import line13 from '../../../assets/lines/line13.png'
import line14 from '../../../assets/lines/line14.png'
import lineA from '../../../assets/lines/lineA.jpg'
import lineB from '../../../assets/lines/lineB.png'


export default {
  line1: {
    query: 'stations-ratp?line=1',
    coordinates: [],
    color: '#F8B900',
    name: 'Métro 1'
  },
  line2: {
    query: 'stations-ratp?line=2',
    coordinates: [],
    color: '#0065AE',
    name: 'Métro 2'
  },
  line3: {
    query: 'stations-ratp?line=3',
    coordinates: [],
    color: '#A19817',
    name: 'Métro 3'
  },
  line3bis: {
    query: 'stations-ratp?line=3bis',
    coordinates: [],
    color: '#8AD4DF',
    name: 'Métro 3bis'
  },
  line4: {
    query: 'stations-ratp?line=4',
    coordinates: [],
    color: '#BE418C',
    name: 'Métro 4'
  },
  line5: {
    query: 'stations-ratp?line=5',
    coordinates: [],
    color: '#F19043',
    name: 'Métro 5'
  },
  line6: {
    query: 'stations-ratp?line=6',
    coordinates: [],
    color: '#84C28E',
    name: 'Métro 6'
  },
  line7: {
    query: 'stations-ratp?line=7',
    coordinates: [],
    color: '#F2A4B7',
    name: 'Métro 7'
  },
  line7bis: {
    query: 'stations-ratp?line=7bis',
    coordinates: [],
    color: '#84C28E',
    name: 'Métro 7bis'
  },
  line8: {
    query: 'stations-ratp?line=8',
    coordinates: [],
    color: '#CDACCF',
    name: 'Métro 8'
  },
  line9: {
    query: 'stations-ratp?line=9',
    coordinates: [],
    color: '#D5C900',
    name: 'Métro 9'
  },
  line10: {
    query: 'stations-ratp?line=10',
    coordinates: [],
    color: '#E4B325',
    name: 'Métro 10'
  },
  line11: {
    query: 'stations-ratp?line=11',
    coordinates: [],
    color: '#8C5E24',
    name: 'Métro 11'
  },
  line12: {
    query: 'stations-ratp?line=12',
    coordinates: [],
    color: '#007E49',
    name: 'Métro 12'
  },
  line13: {
    query: 'stations-ratp?line=13',
    coordinates: [],
    color: '#86CFF3',
    name: 'Métro 13'
  },
  line14: {
    query: 'stations-ratp?line=14',
    coordinates: [],
    color: '#612180',
    name: 'Métro 14'
  },
  lineA: {
    query: 'stations-ratp?line=A',
    coordinates: [],
    color: '#E42212',
    name: 'RER A'
  },
  lineB: {
    query: 'stations-ratp?line=B',
    coordinates: [],
    color: '#5190CD',
    name: 'RER B'
  }
}

export function standardizeLineBranches(lines) {
  const branches = {
    line7brancheI: {
      coordinates: [lines.line7.coordinates[28], lines.line7.coordinates[29], lines.line7.coordinates[30], lines.line7.coordinates[31], lines.line7.coordinates[32], lines.line7.coordinates[33]],
      color: lines.line7.color,
      name: 'Métro 7 (branche Ivry)'
    },
    line7brancheV: {
      coordinates: [lines.line7.coordinates[28], lines.line7.coordinates[34], lines.line7.coordinates[35], lines.line7.coordinates[36], lines.line7.coordinates[37]],
      color: lines.line7.color,
      name: 'Métro 7 (branche Villejuif)',

    },
    line10brancheO: {
      coordinates: [lines.line10.coordinates[0], lines.line10.coordinates[1]],
      color: lines.line10.color,
      name: 'Métro 10 (branche ouest)',

    },
    line10brancheM: {
      coordinates: [lines.line10.coordinates[1], lines.line10.coordinates[2], lines.line10.coordinates[3], lines.line10.coordinates[4], lines.line10.coordinates[5]],
      color: lines.line10.color,
      name: 'Métro 10 (branche Molitor)',

    },
    line10brancheA: {
      coordinates: [lines.line10.coordinates[5], lines.line10.coordinates[6], lines.line10.coordinates[7], lines.line10.coordinates[8], lines.line10.coordinates[1]],
      color: lines.line10.color,
      name: 'Métro 10 (branche Auteuil)',

    },
    line13brancheS: {
      coordinates: [lines.line13.coordinates[0], lines.line13.coordinates[1], lines.line13.coordinates[2], lines.line13.coordinates[3], lines.line13.coordinates[4], lines.line13.coordinates[5], lines.line13.coordinates[6], lines.line13.coordinates[7], lines.line13.coordinates[14]],
      color: lines.line13.color,
      name: 'Métro 13 (branche St-Denis)',

    },
    line13brancheA: {
      coordinates: [lines.line13.coordinates[8], lines.line13.coordinates[9], lines.line13.coordinates[10], lines.line13.coordinates[11], lines.line13.coordinates[12], lines.line13.coordinates[13], lines.line13.coordinates[14]],
      color: lines.line13.color,
      name: 'Métro 13 (branche Asnières)',

    },
    lineAParis: {
      coordinates: [lines.lineA.coordinates[18], lines.lineA.coordinates[19], lines.lineA.coordinates[20], lines.lineA.coordinates[21], lines.lineA.coordinates[22], lines.lineA.coordinates[23], lines.lineA.coordinates[24], lines.lineA.coordinates[25]],
      color: lines.lineA.color,
      name: 'RER A (Paris)',
    },
    lineAbrancheC: {
      coordinates: [lines.lineA.coordinates[0], lines.lineA.coordinates[1], lines.lineA.coordinates[2], lines.lineA.coordinates[3], lines.lineA.coordinates[4], lines.lineA.coordinates[5], lines.lineA.coordinates[6], lines.lineA.coordinates[7], lines.lineA.coordinates[8], lines.lineA.coordinates[18]],
      color: lines.lineA.color,
      name: 'RER A (Branche Cergy)',
    },
    lineAbrancheP: {
      coordinates: [lines.lineA.coordinates[9], lines.lineA.coordinates[10], lines.lineA.coordinates[6]],
      color: lines.lineA.color,
      name: 'RER A (Branche Poissy)',
    },
    lineAbrancheS: {
      coordinates: [lines.lineA.coordinates[11], lines.lineA.coordinates[12], lines.lineA.coordinates[13], lines.lineA.coordinates[14], lines.lineA.coordinates[15], lines.lineA.coordinates[16], lines.lineA.coordinates[17], lines.lineA.coordinates[18]],
      color: lines.lineA.color,
      name: 'RER A (Branche Saint-Germain-en-Laye)',
    },
    lineAbrancheM: {
      coordinates: [lines.lineA.coordinates[25], lines.lineA.coordinates[26], lines.lineA.coordinates[27], lines.lineA.coordinates[28], lines.lineA.coordinates[29], lines.lineA.coordinates[30], lines.lineA.coordinates[31], lines.lineA.coordinates[32], lines.lineA.coordinates[33], lines.lineA.coordinates[34], lines.lineA.coordinates[35], lines.lineA.coordinates[36]],
      color: lines.lineA.color,
      name: 'RER A (Branche Marne-la-Vallée)',
    },
    lineAbrancheB: {
      coordinates: [lines.lineA.coordinates[25], ...lines.lineA.coordinates.slice(37)],
      color: lines.lineA.color,
      name: 'RER A (Branche Boissy-St-Léger)',
    },
    lineBParis: {
      coordinates: [lines.lineB.coordinates[9], lines.lineB.coordinates[10], lines.lineB.coordinates[11], lines.lineB.coordinates[12], lines.lineB.coordinates[13], lines.lineB.coordinates[14], lines.lineB.coordinates[15], lines.lineB.coordinates[16], lines.lineB.coordinates[17], lines.lineB.coordinates[18], lines.lineB.coordinates[19], lines.lineB.coordinates[20], lines.lineB.coordinates[21], lines.lineB.coordinates[22], lines.lineB.coordinates[23], lines.lineB.coordinates[24], lines.lineB.coordinates[25], lines.lineB.coordinates[26]],
      color: lines.lineB.color,
      name: 'RER B (Paris)',
    },
    lineBbrancheM: {
      coordinates: [lines.lineB.coordinates[0], lines.lineB.coordinates[1], lines.lineB.coordinates[2], lines.lineB.coordinates[3], lines.lineB.coordinates[9]],
      color: lines.lineB.color,
      name: 'RER B (Branche Mitry-Claye)',
    },
    lineBbrancheC: {
      coordinates: [lines.lineB.coordinates[4], lines.lineB.coordinates[5], lines.lineB.coordinates[6], lines.lineB.coordinates[7], lines.lineB.coordinates[8], lines.lineB.coordinates[9]],
      color: lines.lineB.color,
      name: 'RER B (Branche CDG)',
    },
    lineBbrancheS: {
      coordinates: [lines.lineB.coordinates[26], lines.lineB.coordinates[27], lines.lineB.coordinates[28], lines.lineB.coordinates[29], lines.lineB.coordinates[30], lines.lineB.coordinates[31], lines.lineB.coordinates[32], lines.lineB.coordinates[33], lines.lineB.coordinates[34], lines.lineB.coordinates[35], lines.lineB.coordinates[36], lines.lineB.coordinates[37], lines.lineB.coordinates[38], lines.lineB.coordinates[39], lines.lineB.coordinates[40], lines.lineB.coordinates[41], lines.lineB.coordinates[42], lines.lineB.coordinates[43], lines.lineB.coordinates[44]],
      color: lines.lineB.color,
      name: 'RER B (Branche Saint-Rémy-Lès-Chevreuse)',
    },
    lineBbrancheR: {
      coordinates: [lines.lineB.coordinates[26], ...lines.lineB.coordinates.slice(45)],
      color: lines.lineB.color,
      name: 'RER B (Branche Robinson)',
    },
  }
  lines.lineA.coordinates = []
  lines.lineB.coordinates = []
  lines.line7.nbStations = lines.line7.coordinates.length
  lines.line7.coordinates = lines.line7.coordinates.slice(0, -9)
  lines.line7bis.nbStations = lines.line7bis.coordinates.length
  lines.line7bis.coordinates = [...lines.line7bis.coordinates, lines.line7bis.coordinates[4]]
  lines.line10.nbStations = lines.line10.coordinates.length
  lines.line10.coordinates = [lines.line10.coordinates[5], ...lines.line10.coordinates.slice(9)]
  lines.line13.nbStations = lines.line13.coordinates.length
  lines.line13.coordinates = lines.line13.coordinates.slice(14)
  lines = {...lines, ...branches}
  return lines
}

export function getLineLogo(line) {
  let logo = ''
  switch (line) {
    case 'A':
      logo = lineA
      break
    case 'B':
      logo = lineB
      break
    case '1':
      logo = line1
      break
    case '2':
      logo = line2
      break
    case '3':
      logo = line3
      break
    case '3bis':
      logo = line3bis
      break
    case '4':
      logo = line4
      break
    case '5':
      logo = line5
      break
    case '6':
      logo = line6
      break
    case '7':
      logo = line7
      break
    case '7bis':
      logo = line7bis
      break
    case '8':
      logo = line8
      break
    case '9':
      logo = line9
      break
    case '10':
      logo = line10
      break
    case '11':
      logo = line11
      break
    case '12':
      logo = line12
      break
    case '13':
      logo = line13
      break
    case '14':
      logo = line14
      break
    default:
  }
  return logo
}
