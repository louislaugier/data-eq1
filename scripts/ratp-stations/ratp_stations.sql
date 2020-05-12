CREATE TABLE ratp_stations (
    id INTEGER NOT NULL,
    name VARCHAR NOT NULL,
    address VARCHAR NOT NULL,
    coordinates VARCHAR NOT NULL
);

-- Run main.go

ALTER TABLE ratp_stations
    ADD train_outside BOOLEAN DEFAULT FALSE NOT NULL,
    ADD transport_type VARCHAR DEFAULT 'bus' NOT NULL,
    ADD line VARCHAR NULL,
    ADD line_index INTEGER;
-- Manually enter line index for each station to trace lines on map

-- Define records with non-UPPERCASE names as metros
UPDATE ratp_stations
    SET transport_type = 'métro'
    WHERE name <> UPPER(name);

-- Define RER stations with unique names
UPDATE ratp_stations SET transport_type = 'RER' WHERE "name" IN ('Achères Grand Cormier','Achères-Ville','Aéroport Charles de Gaulle 1','Aéroport Charles de Gaulle 2 TGV','Antony','Arcueil-Cachan','Auber','Aulnay-sous-Bois','Bagneux','Boissy-Saint-Léger','Bourg-la-Reine','Bry-sur-Marne','Bures-sur-Yvette','Bussy-Saint-Georges','Cergy-Le-Haut','Cergy-Préfecture','Cergy-Saint-Christophe','Champigny','Châtelet-Les Halles','Chatou-Croissy','Cité Universitaire','Conflans-Fin d''Oise','Courcelle-sur-Yvette','Drancy','Fontaine-Michalon','Fontenay-aux-Roses','Fontenay-sous-Bois','Gentilly','Gif-sur-Yvette','Houilles Carrières-sur-Seine','Joinville-le-Pont','La Courneuve-Aubervilliers','La Croix-de-Berny','La Hacquinière','La Plaine-Stade de France','La Varenne-Chennevières','Laplace','Le Blanc-Mesnil','Le Bourget','Le Guichet','Le Parc de Saint-Maur','Le Vésinet-Centre','Le Vésinet-Le Pecq','Les Baconnets','Lognes','Lozère','Luxembourg','Maisons-Laffitte','Marne-la-Vallée Chessy','Massy-Palaiseau','Massy-Verrières','Mitry-Claye','Nanterre-Préfecture','Nanterre-Université','Nanterre-Ville','Neuilly-Plaisance','Neuville-Université','Nogent-sur-Marne','Noisiel','Noisy-Champs','Noisy-le-Grand (Mont d''Est)','Orsay-Ville','Palaiseau','Palaiseau Villebon','Parc de Sceaux','Parc des Expositions','Poissy','Port-Royal','Robinson','Rueil-Malmaison','Saint-Germain-en-Laye','Saint-Maur Créteil','Saint-Michel Notre-Dame','Saint-Rémy-lès-Chevreuse','Sartrouville','Sceaux','Sevran-Beaudottes','Sevran Livry','Sucy Bonneuil','Torcy','Val d''Europe','Val de Fontenay','Vert-Galant','Villeparisis','Villepinte','Vincennes');

-- Define RER stations which have metro stations with the same name by their coordinates
-- Charles de Gaulle - Étoile
UPDATE ratp_stations SET transport_type = 'RER' WHERE 'coordinates' = '48.8739040986,2.29503190105';
-- Gare de Lyon
UPDATE ratp_stations SET transport_type = 'RER' WHERE 'coordinates' = '48.844472156,2.3739114213';
-- Gare du Nord
UPDATE ratp_stations SET transport_type = 'RER' WHERE 'coordinates' = '48.8808392058,2.35657701535';
-- La Défense - Grande Arche
UPDATE ratp_stations SET transport_type = 'RER' WHERE 'coordinates' = '48.8918267548,2.23799204322';
-- Nation
UPDATE ratp_stations SET transport_type = 'RER' WHERE 'coordinates' = '48.8484448201,2.39590723714';
-- Denfert-Rochereau
UPDATE ratp_stations SET transport_type = 'RER' WHERE 'coordinates' = '48.8484448201,2.39590723714';

-- Define exterior train stations which have only 1 line
UPDATE ratp_stations SET train_outside = TRUE WHERE "name" IN ('Bir-Hakeim (Grenelle)','Cambronne','Châtillon Montrouge','Chevaleret','Corvisart','Dupleix','Glacière','La Chapelle','Malakoff-Rue Etienne Dolet','Nationale','Quai de la Gare','Sèvres-Lecourbe','Bel-Air','Créteil-L''Echat (Hôpital Henri Mondor)','Créteil-Préfecture (Hôtel de Ville)','Créteil-Université','Passy','Pointe du Lac','Quai de la Rapée','Saint-Jacques','Cergy-Le-Haut','Cergy-Saint-Christophe','Fontenay-sous-Bois','Marne-la-Vallée Chessy','Neuville-Université','Achères-Ville','Poissy','Achères Grand Cormier','Maisons-Laffitte','Sartrouville','Houilles Carrières-sur-Seine','Le Vésinet-Le Pecq','Le Vésinet-Centre','Chatou-Croissy','Nanterre-Ville','Nanterre-Université','Le Parc de Saint-Maur','La Varenne-Chennevières','Sucy Bonneuil','Boissy-Saint-Léger','Noisy-Champs','Noisiel','Lognes','Torcy','Val d''Europe','Conflans-Fin d''Oise','Rueil-Malmaison','Joinville-le-Pont','Saint-Maur Créteil','Champigny','Neuilly-Plaisance','Bry-sur-Marne','Aéroport Charles de Gaulle 2 TGV','Antony','Arcueil-Cachan','Aulnay-sous-Bois','Bagneux','Bourg-la-Reine','Bures-sur-Yvette','Cité Universitaire','Courcelle-sur-Yvette','Drancy','Fontaine-Michalon','Fontenay-aux-Roses','Gentilly','Gif-sur-Yvette','La Courneuve-Aubervilliers','La Croix-de-Berny','La Hacquinière','La Plaine-Stade de France','Laplace','Le Blanc-Mesnil','Le Bourget','Le Guichet','Les Baconnets','Lozère','Massy-Palaiseau','Massy-Verrières','Mitry-Claye','Orsay-Ville','Palaiseau','Palaiseau Villebon','Parc de Sceaux','Parc des Expositions','Port-Royal','Robinson','Saint-Rémy-lès-Chevreuse','Sceaux','Sevran Livry','Vert-Galant','Villeparisis','Villepinte');

-- Define exterior train stations which have several lines by their coordinates
-- Bastille, Barbès-Rochechouart, Stalingrad, Jaurès, Gare d'Austerlitz, La Motte-Picquet-Grenelle, Denfert-Rochereau
UPDATE ratp_stations SET train_outside = TRUE WHERE "coordinates" IN ('48.8533891123,2.36916460221','48.8839317699,2.3493557086','48.8841892942,2.36703044197','48.8827685467,2.36994579899','48.8437198614,2.36391621716','48.8490189283,2.29776216071');

-- Add lines for train stations which have only 1 line
UPDATE ratp_stations SET line = '1' WHERE "name" IN ('Argentine','Bérault','Château de Vincennes','Esplanade de la Défense','George V','La Défense (Grande Arche)','Les Sablons (Jardin d''acclimatation)','Louvre-Rivoli','Pont de Neuilly','Porte de Vincennes','Porte Maillot','Saint-Mandé','Saint-Paul (Le Marais)','Tuileries');
UPDATE ratp_stations SET line = '2' WHERE "name" IN ('La Chapelle','Alexandre-Dumas','Anvers','Avron','Blanche','Colonel Fabien','Courcelles','Couronnes','Ménilmontant','Monceau','Philippe Auguste','Porte Dauphine (Maréchal de Lattre de Tassigny)','Rome','Ternes','Victor Hugo');
UPDATE ratp_stations SET line = '3' WHERE "name" IN ('Anatole-France','Bourse','Europe','Gallieni (Parc de Bagnolet)','Louise Michel','Malesherbes','Parmentier','Pereire','Pont de Levallois-Bécon','Porte de Bagnolet','Porte de Champerret','Quatre Septembre','Rue Saint-Maur','Sentier','Temple','Wagram');
UPDATE ratp_stations SET line = '3bis' WHERE "name" IN ('Pelleport','Saint-Fargeau');
UPDATE ratp_stations SET line = '4' WHERE "name" IN ('Alésia','Château d''Eau','Château Rouge','Cité','Etienne Marcel','Les Halles','Mairie de Montrouge','Mouton-Duvernet','Porte de Clignancourt','Porte d''Orléans (Général Leclerc)','Saint-Germain des Prés','Saint-Michel','Saint-Placide','Saint-Sulpice','Simplon','Vavin');
UPDATE ratp_stations SET line = '5' WHERE "name" IN ('Quai de la Rapée','Bobigny-Pablo-Picasso','Bobigny-Pantin (Raymond Queneau)','Bréguet-Sabin','Campo-Formio','Eglise de Pantin','Hoche','Jacques-Bonsergent','Laumière','Ourcq','Porte de Pantin','Richard-Lenoir','Saint-Marcel');
UPDATE ratp_stations SET line = '6' WHERE "name" IN ('Bir-Hakeim (Grenelle)','Cambronne','Chevaleret','Corvisart','Dupleix','Glacière','Quai de la Gare','Sèvres-Lecourbe','Bel-Air','Passy','Saint-Jacques','Boissière','Dugommier','Edgar-Quinet','Kléber','Picpus','Nationale');
UPDATE ratp_stations SET line = '7' WHERE "name" IN ('Aubervilliers Pantin (4 Chemins)','Cadet','Censier-Daubenton','Château Landon','Corentin-Cariou','Crimée','Fort d''Aubervilliers','La Courneuve-8-Mai-1945','Le Kremlin-Bicêtre','Le Peletier','Les Gobelins','Mairie d''Ivry','Maison Blanche','Pierre et Marie Curie','Place Monge (Jardin des Plantes)','Poissonnière','Pont Marie (Cité des Arts)','Pont Neuf','Porte de Choisy','Porte de la Villette','Porte d''Italie','Porte d''Ivry','Riquet','Sully-Morland','Tolbiac','Villejuif-Léo Lagrange','Villejuif-Louis Aragon','Villejuif-Paul Vaillant Couturier (Hôpital Paul Brousse)');
UPDATE ratp_stations SET line = '7bis' WHERE "name" IN ('Bolivar','Botzaris','Buttes-Chaumont','Danube','Pré-Saint-Gervais');
UPDATE ratp_stations SET line = '8' WHERE "name" IN ('Créteil-L''Echat (Hôpital Henri Mondor)','Créteil-Préfecture (Hôtel de Ville)','Créteil-Université','Pointe du Lac','Balard','Boucicaut','Charenton-Ecoles','Chemin Vert','Commerce','Ecole Militaire','Ecole Vétérinaire de Maisons-Alfort','Faidherbe-Chaligny','Félix Faure','Filles du Calvaire','La Tour-Maubourg','Ledru-Rollin','Liberté','Lourmel','Maisons-Alfort-Les Juilliottes','Maisons-Alfort-Stade','Michel Bizot','Montgallet','Porte de Charenton','Porte Dorée','Saint-Sébastien-Froissart');
UPDATE ratp_stations SET line = '9' WHERE "name" IN ('Alma-Marceau','Billancourt','Buzenval','Charonne','Croix-de-Chavaux (Jacques Duclos)','Exelmans','Iéna','Jasmin','La Muette','Mairie de Montreuil','Maraîchers','Marcel Sembat','Pont de Sèvres','Porte de Montreuil','Porte de Saint-Cloud','Ranelagh','Robespierre','Rue de la Pompe (Avenue Georges Mandel)','Saint-Ambroise','Saint-Augustin','Saint-Philippe du Roule','Voltaire (Léon Blum)','Rue des Boulets');
UPDATE ratp_stations SET line = '10' WHERE "name" IN ('Avenue Emile-Zola','Boulogne-Jean-Jaurès','Boulogne Pont de Saint-Cloud','Cardinal-Lemoine','Chardon-Lagache','Charles Michels','Cluny-La Sorbonne','Eglise d''Auteuil','Javel-André-Citroen','Mabillon','Maubert-Mutualité','Mirabeau','Porte d''Auteuil','Ségur','Vaneau');
UPDATE ratp_stations SET line = '11' WHERE "name" IN ('Goncourt (Hôpital Saint-Louis)','Jourdain','Mairie des Lilas','Pyrénées','Rambuteau','Télégraphe');
UPDATE ratp_stations SET line = '12' WHERE "name" IN ('Abbesses','Assemblée Nationale','Convention','Corentin-Celton','Falguière','Front Populaire','Jules Joffrin','Lamarck-Caulaincourt','Mairie d''Issy','Marx-Dormoy','Notre-Dame de Lorette','Notre-Dame des Champs','Porte de la Chapelle','Porte de Versailles','Rennes','Rue du Bac','Saint-Georges','Solférino','Trinité-d''Estienne d''Orves','Vaugirard (Adolphe Chérioux)','Volontaires');
UPDATE ratp_stations SET line = '13' WHERE "name" IN ('Asnières-Gennevilliers Les Courtilles','Châtillon Montrouge','Malakoff-Rue Etienne Dolet','Basilique de Saint-Denis','Brochant','Carrefour-Pleyel','Gabriel-Péri','Gaîté','Garibaldi','Guy-Môquet','La Fourche','Les Agnettes','Les Courtilles','Liège','Mairie de Clichy','Mairie de Saint-Ouen','Malakoff-Plateau de Vanves','Pernety','Plaisance','Porte de Clichy','Porte de Saint-Ouen','Porte de Vanves','Saint-Denis - Porte de Paris','Saint-Denis-Université','Saint-François-Xavier','Varenne');
UPDATE ratp_stations SET line = '14' WHERE "name" IN ('Bibliothèque-François Mitterrand','Cour Saint-Emilion','Olympiades');
UPDATE ratp_stations SET line = 'A' WHERE "name" IN ('Cergy-Le-Haut','Cergy-Saint-Christophe','Fontenay-sous-Bois','Marne-la-Vallée Chessy','Cergy-Préfecture','Saint-Germain-en-Laye','Nanterre-Préfecture','Vincennes','Nogent-sur-Marne','Noisy-le-Grand (Mont d''Est)','Bussy-Saint-Georges','Neuville-Université','Achères-Ville','Poissy','Achères Grand Cormier','Maisons-Laffitte','Sartrouville','Houilles Carrières-sur-Seine','Le Vésinet-Le Pecq','Le Vésinet-Centre','Chatou-Croissy','Nanterre-Ville','Nanterre-Université','Le Parc de Saint-Maur','La Varenne-Chennevières','Sucy Bonneuil','Boissy-Saint-Léger','Noisy-Champs','Noisiel','Lognes','Torcy','Val d''Europe','Conflans-Fin d''Oise','Rueil-Malmaison','Joinville-le-Pont','Saint-Maur Créteil','Champigny','Neuilly-Plaisance','Bry-sur-Marne','Val de Fontenay');
UPDATE ratp_stations SET line = 'B' WHERE "name" IN ('Aéroport Charles de Gaulle 1','Aéroport Charles de Gaulle 2 TGV','Antony','Arcueil-Cachan','Aulnay-sous-Bois','Bagneux','Bourg-la-Reine','Bures-sur-Yvette','Cité Universitaire','Courcelle-sur-Yvette','Drancy','Fontaine-Michalon','Fontenay-aux-Roses','Gentilly','Gif-sur-Yvette','La Courneuve-Aubervilliers','La Croix-de-Berny','La Hacquinière','La Plaine-Stade de France','Laplace','Blanc-Mesnil','Le Bourget','Le Guichet','Les Baconnets','Lozère','Luxembourg','Massy-Palaiseau','Massy-Verrières','Mitry-Claye','Orsay-Ville','Palaiseau','Palaiseau Villebon','Parc de Sceaux','Parc des Expositions','Port-Royal','Robinson','Saint-Rémy-lès-Chevreuse','Sceaux','Sevran-Beaudottes','Sevran Livry','Vert-Galant','Villeparisis','Villepinte');

-- Add lines for train stations which have several lines by their coordinates
-- Bastille, Champs-Élysées - Clémenceau, Charles de Gaulle - Étoile, Châtelet, Concorde, Franklin D. Roosevelt, Gare de Lyon, Hôtel de Ville, Nation, Palais-Royal (Musée du Louvre), Reuilly-Diderot
UPDATE ratp_stations SET line = '1' WHERE "coordinates" IN ('48.8533891123,2.36916460221','48.8673665705,2.31419106339','48.8739310943,2.29512725165','48.8589561813,2.34736125521','48.8656780964,2.3211937718','48.8690104279,2.3102531809','48.8455597871,2.37344920496','48.8573559272,2.35207356686','48.8494239009,2.3970250383','48.8633873851,2.33548383677','48.8472175011,2.38720417721');
-- Barbès-Rochechouart, Stalingrad, Jaurès, Belleville, Charles de Gaulle - Étoile, Nation, Père-Lachaise, Pigalle, Place de Clichy, Villiers
UPDATE ratp_stations SET line = '2' WHERE "coordinates" IN ('48.8839317699,2.3493557086','48.8841892942,2.36703044197','48.8827685467,2.36994579899','48.8722873206,2.37673767732','48.8737152683,2.29475956285','48.8490101536,2.39766461279','48.8631612051,2.38726040018','48.8824225284,2.3372546821','48.8834556516,2.32737499341','48.8813242939,2.31659676151');
-- Arts et Métiers, Gambetta, Havre - Caumartin, Opéra, Père-Lachaise, Réaumur-Sébastopol, République, Saint-Lazare, Villiers
UPDATE ratp_stations SET line = '3' WHERE "coordinates" IN ('48.8655518636,2.35610815512','48.8647827007,2.39844583377','48.8734618907,2.32848048447','48.8707209763,2.33225474357','48.8629456617,2.38689240454','48.8662803586,2.35248479684','48.8672579436,2.36403730364','48.8755376704,2.32519648998','48.8814051004,2.31614705091');
-- Gambetta, Porte des Lilas
UPDATE ratp_stations SET line = '3bis' WHERE "coordinates" IN ('48.865034298,2.39854148691','48.8769382235,2.40637656693');
-- Barbès-Rochechouart, Châtelet, Gare de l'Est, Denfert-Rochereau, Gare du Nord, Marcadet - Poissonniers, Montparnasse - Bienvenüe, Odéon, Raspail, Réaumur-Sébastopol, Strasbourg-Saint-Denis
UPDATE ratp_stations SET line = '4' WHERE "coordinates" IN ('48.8836801093,2.34953280365','48.858300056,2.34786508758','48.8764170529,2.35898704646','48.8333151859,2.33344276383','48.8793833026,2.35637208136','48.8915439306,2.34937111538','48.8441895085,2.32444123243','48.8521980452,2.33878000304','48.8388783786,2.33067841929','48.8660466696,2.35262095402','48.869227931,2.35452917539');
-- Bastille, Stalingrad, Jaurès, Gare d'Austerlitz, Gare de l'Est, Gare du Nord, Oberkampf, Place d'Italie, République
UPDATE ratp_stations SET line = '5' WHERE "coordinates" IN ('48.8529756747,2.36921882444','48.8841532631,2.36735747747','48.8823280249,2.37047696877','48.8437198614,2.36391621716','48.8768126254,2.35825144008','48.8797430642,2.35460081047','48.8649742906,2.36763241156','48.8315523026,2.35533344322','48.8674108288,2.36358782556');
-- La Motte-Picquet-Grenelle, Bercy, Charles de Gaulle - Étoile, Daumesnil, Denfert-Rochereau, Montparnasse - Bienvenüe, Nation, Place d'Italie, Raspail, Trocadéro, Pasteur
UPDATE ratp_stations SET line = '6' WHERE "coordinates" IN ('48.8490189283,2.29776216071','48.840542783,2.37940946312','48.8739938407,2.29466396183','48.8396821402,2.39558417035','48.8339982367,2.33323850116','48.8440276362,2.32356979722','48.8474552623,2.39777174361','48.8312826555,2.35545587036','48.8391569815,2.33047415116','48.863089135,2.28684037849','48.8420128041,2.31329004384');
-- Stalingrad, Châtelet, Chaussée d'Antin - La Fayette, Gare de l'Est, Jussieu, Louis Blanc, Opéra, Palais-Royal (Musée du Louvre), Place d'Italie, Pyramides
UPDATE ratp_stations SET line = '7' WHERE "coordinates" IN ('48.8840719864,2.36887006278','48.8572036455,2.34722471106','48.8728599894,2.33304475093', '48.876695836,2.35799251055','48.8459684322,2.35480739269','48.8809004282,2.36510721826','48.8706310801,2.33173704836','48.8623718215,2.33657360074','48.8311118383,2.35579614571','48.866496987,2.33409421526');
-- Jaurès, Louis Blanc, Place des Fêtes
UPDATE ratp_stations SET line = '7bis' WHERE "coordinates" IN ('48.8828762162,2.37058633617','48.8814934289,2.36577524235','48.8767294689,2.39314603177');
-- Bastille, La Motte-Picquet-Grenelle, Bonne-Nouvelle, Concorde, Daumesnil, Grands Boulevards, Invalides, Madeleine, Opéra, République, Reuilly-Diderot, Richelieu-Drouot, Strasbourg-Saint-Denis
UPDATE ratp_stations SET line = '8' WHERE "coordinates" IN ('48.853739623,2.36916481786','48.8488212667,2.29795296403','48.8705767536,2.34849429371','48.8655072215,2.32036285586','48.8393764631,2.39578806042','48.8715118171,2.34320837754','48.8610934676,2.31464335535','48.8695789706,2.32416248879','48.8705592165,2.33255448205','48.867671519,2.36332912078','48.847352877,2.38584254004','48.8720961238,2.33932555586','48.8696323652,2.35450206309');
-- Bonne-Nouvelle, Chaussée d'Antin - La Fayette, Franklin D. Roosevelt, Grands Boulevards, Havre - Caumartin, Michel-Ange - Auteuil, Michel-Ange - Molitor, Miromesnil, Nation, Oberkampf, République, Richelieu-Drouot, Strasbourg-Saint-Denis, Trocadéro
UPDATE ratp_stations SET line = '9' WHERE "coordinates" IN ('48.8704958323,2.34886211744','48.8729858247,2.33342622405','48.8688126319,2.30992632447','48.8715747436,2.34289503198','48.8734259765,2.3289709733','48.8479245928,2.26423590578','48.845056129,2.26191142328','48.8734603255,2.31604128892','48.8481112316,2.3980040128','48.8649024159,2.36753701348','48.8676174865,2.36381952143','48.872006239,2.33991139061','48.8693807801,2.3540796522','48.8626669059,2.28726306155');
-- Gare d'Austerlitz, La Motte-Picquet-Grenelle, Duroc, Jussieu, Michel-Ange - Auteuil, Michel-Ange - Molitor, Odéon, Sèvres-Babylone
UPDATE ratp_stations SET line = '10' WHERE "coordinates" IN ('48.8434142241,2.36418839287','48.848803206,2.29769423466','48.8468487689,2.31693731083','48.8457617573,2.35454859491','48.8479339461,2.26483507717','48.8449844867,2.26232003904','48.8523148812,2.33886172034','48.8509392884,2.32611465482');
-- Arts et métiers, Belleville, Châtelet, Hôtel de Ville, Place des Fêtes, Porte des Lilas, République
UPDATE ratp_stations SET line = '11' WHERE "coordinates" IN ('48.865381176,2.35564493121','48.871945798,2.37675104137','48.8588572844,2.34776985674','48.8578053731,2.35146077651','48.8772239477,2.39276505025','48.8772347167,2.40652684305','48.8672401854,2.36305644595');
-- Concorde, Madeleine, Marcadet - Poissonniers, Montparnasse - Bienvenüe, Pigalle, Saint-Lazare, Sèvres-Babylone, Pasteur
UPDATE ratp_stations SET line = '12' WHERE "coordinates" IN ('48.8654893909,2.32141178921','48.8695071276,2.32468019959','48.8916068227,2.34954831099','48.8420770202,2.32101044956','48.8820270896,2.33721380175','48.8754209778,2.32669527207','48.851577469,2.32689078834','48.8428305421,2.31266330776');
-- Champs-Élysées Clémenceau, Duroc, Invalides, Miromesnil, Montparnasse - Bienvenüe, Place de Clichy, Saint-Lazare
UPDATE ratp_stations SET line = '13' WHERE "coordinates" IN ('48.8677440261,2.31412278037','48.8469925009,2.31654234477','48.8608507897,2.31454811263','48.8735321406,2.31557802487','48.8418703363,2.321228374','48.8832039999,2.32726602464','48.8756815064,2.32559158482');
-- Bercy, Châtelet, Gare de Lyon, Madeleine, Pyramides, Saint-Lazare
UPDATE ratp_stations SET line = '14' WHERE "coordinates" IN ('48.8401922444,2.37947725727','48.8585696725,2.34793324584','48.844652151,2.37310814755','48.8697947152,2.32461201202','48.8666138306,2.33440753095','48.8760678144,2.32418810077');
-- La Défense, Charles de Gaulle - Étoile, Auber, Nation, Châtelet - Les Halles, Gare de Lyon
UPDATE ratp_stations SET line = 'A' WHERE "coordinates" IN ('48.8918267548,2.23799204322','48.8739040986,2.29503190105','48.8726081799,2.32970681645','48.8484448201,2.39590723714','48.8614637106,2.3468441316','48.844472156,2.3739114213');
-- Châtelet - Les Halles, Denfert-Rochereau, Gare du Nord, Saint-Michel Notre-Dame
UPDATE ratp_stations SET line = 'B' WHERE "coordinates" IN ('48.8617962475,2.34680332871','48.8342768128,2.3322037973','48.8808392058,2.35657701535','48.8536807808,2.34424139962');