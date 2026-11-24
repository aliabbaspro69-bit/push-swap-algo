# Mini TP 1 : Explorer la mini API avec Swagger

## 1-2. Lancement et tests
Fork de louvelmathieu/go-web-2026, lancé avec `go run ./api`.
Tous les endpoints testés via http://localhost:8081/api/docs (Try it out).

## 3. Codes d'erreur
| URL | Code | Réponse |
|---|---|---|
| /api/games/999 | 404 Not Found | `{"error": "jeu introuvable"}` |
| /api/games/abc | 400 Bad Request | `{"error": "l'identifiant doit être un nombre"}` |

Remarque : Swagger bloque `abc` côté navigateur (validation "Value must be an integer"),
j'ai donc testé directement dans la barre d'adresse.

## 4. Jeu 4 : Street Fighter 6
Le jeu 4 a **4 tournois** (`/api/games/4/tournaments`) :
- Fight Night Marseille (Marseille)
- Combo Breaker EU (Paris)
- Rising Dragon (Tokyo)
- Hadoken Series (Bruxelles)

## 5. Filtres sur /api/games
- `?genre=course` : Rocket League (2015), Trackmania (2020), Mario Kart 8 Deluxe (2017)
- `?year=2018` : Celeste, Dead Cells

## 6. /api/stats (authentification)
| Situation | Code | Réponse |
|---|---|---|
| Sans clé | 401 | clé absente ou inconnue : ajoutez l'en-tête X-API-Key |
| Clé `eleve` | 403 | clé reconnue, mais pas le droit de voir les statistiques |
| Clé `prof` | 200 | `{"games": 12, "tournaments": 20}` |

401 = pas identifié, 403 = identifié mais pas autorisé.

## Bonus
| Endpoint | Code | Détail |
|---|---|---|
| /api/broken | 500 Internal Server Error | `{"error": "panne simulée"}` |
| /api/slow | 200 OK | réponse correcte (12 jeux) mais lente : X secondes |
