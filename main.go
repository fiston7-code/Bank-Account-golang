package main

import "fmt"

type Account struct {
	Balance  int
	Loan     int
	IsActive bool
}

// open account methode

func (a *Account) OpenAccount() {
	if a.IsActive {

		fmt.Println("le compte est deja ouvert !")
		return

	}
	a.IsActive = true
	a.Balance = 500
	fmt.Printf("felicitation votre solde est de : %d\n", a.Balance)

}

// deposit methode

func (a *Account) Deposit(amount int) {
	if !a.IsActive {
		fmt.Println("erreur votre compte ne pas actif !")
		return
	}
	a.Balance += amount
}

// Withdraw methode

func (a *Account) Withdraw(amount int) {
	if !a.IsActive {
		fmt.Println("desole votre compte n'est pas actif")
		return
	}
	if a.Balance < amount {
		fmt.Println("desole le solde de votre compte est insuffisant pour retirer ce motant")
		return

	}
	a.Balance -= amount

}

//     borrow methode

func (a *Account) Borrow(amount int) {
	if !a.IsActive || a.Loan > 0 {
		fmt.Println("Erreur : Compte inactif ou dette déjà existante.")
		return
	} else {
		a.Loan = amount
		a.Balance += amount

		fmt.Printf("Prêt accordé ! +%d$ sur votre solde. Dette actuelle : %d$\n", amount, a.Loan)

	}

}

// // repay methode
func (a *Account) Repay(amount int) bool { // <-- On ajoute bool ici
	if !a.IsActive || amount > a.Loan || a.Balance < amount {
		// On garde tes messages d'erreur à l'intérieur
		if !a.IsActive {
			fmt.Println("Compte inactif.")
		}
		if amount > a.Loan {
			fmt.Printf("Vous ne devez que %d$...\n", a.Loan)
		}
		if a.Balance < amount {
			fmt.Println("Solde insuffisant.")
		}

		return false // <-- On renvoie "Faux" si une sécurité bloque
	}

	a.Balance -= amount
	a.Loan -= amount
	fmt.Printf("Merci ! Remboursement de %d$ effectué.\n", amount)

	return true // <-- On renvoie "Vrai" si tout s'est bien passé
}

//    Fermer() :

//         Vérifier : est-ce que le compte est actif ET le solde est à 0 ET la dette est à 0 ?

//         Si oui -> EstActif = false.

//         Sinon -> afficher pourquoi on ne peut pas fermer (ex: "Il reste de l'argent" ou "Dette en cours").

// close an account
func (a *Account) Close() bool { // Majuscule pour être exportable
	if !a.IsActive {
		fmt.Println("Le compte est déjà fermé.")
		return false
	}

	if a.Balance > 0 || a.Loan > 0 {
		fmt.Printf("Impossible de fermer : Solde (%d$) ou Dette (%d$) non nuls.\n", a.Balance, a.Loan)
		return false
	}

	a.IsActive = false
	fmt.Println("Compte clôturé avec succès. Merci de votre fidélité !")
	return true
}

func main() {

	depot := 0
	retrait := 500
	emprunter := 0
	payer := 0

	account := Account{}
	fmt.Printf("voici le user account initial %d\n", account.Balance)

	account.OpenAccount()
	fmt.Printf("votre compte est actif %t, et votre solde est de: %d\n", account.IsActive, account.Balance)

	account.Deposit(depot)
	fmt.Printf("cher client vous venez de deposer %d$ dans votre compte et votre solde total est de : %d$\n", depot, account.Balance)

	account.Withdraw(retrait)
	fmt.Printf("cher client vous venez de retirer %d$ dans votre compte et le solde restant est de : %d\n", retrait, account.Balance)
	account.Borrow(emprunter)
	fmt.Printf("cher client vous venez de recevoir %d$ et votre solde est de : %d\n", emprunter, account.Balance)
	succes := account.Repay(payer)

	if succes {
		fmt.Printf("Confirmation : Votre nouveau solde est de : %d\n", account.Balance)
	} else {
		fmt.Println("⚠️ Opération annulée par la banque.")
	}

	reussite := account.Close()

	if reussite {
		fmt.Println("Statut final : Fermé (Succès)")
	} else if !account.IsActive {
		fmt.Println("Statut final : Fermé (Déjà fait)")
	} else {
		fmt.Println("Statut final : Toujours ouvert (Erreur)")
	}

}

// Challenge : Système de Gestion Bancaire "ZandoBank"
// 1. La Structure (Le State)

//     Définir une structure Compte avec :

//         Solde (nombre entier)

//         Dette (nombre entier)

//         EstActif (booléen)

// 2. Les Comportements (Les Méthodes)

// Tu dois créer les méthodes suivantes en utilisant des pointeurs :

//     Ouvrir() :

//         Si déjà actif -> ne rien faire.

//         Sinon -> Activer et mettre 500 dans le solde.

//     Deposer(montant) :

//         Si pas actif -> afficher une erreur.

//         Sinon -> Ajouter le montant au solde.

//     Retirer(montant) :

//         Si pas actif -> afficher une erreur.

//         Si le solde est insuffisant -> afficher "Fonds insuffisants".

//         Sinon -> Soustraire du solde.

//     Emprunter(montant) :

//         Si pas actif ou si Dette > 0 -> refuser.

//         Sinon -> Mettre à jour la dette et ajouter le montant au solde.

//     Rembourser() :

//         Si le solde est suffisant pour couvrir TOUTE la dette :

//             Soustraire la dette du solde.

//             Remettre la dette à 0.

//         Sinon -> afficher "Pas assez d'argent pour rembourser".

//     Fermer() :

//         Vérifier : est-ce que le compte est actif ET le solde est à 0 ET la dette est à 0 ?

//         Si oui -> EstActif = false.

//         Sinon -> afficher pourquoi on ne peut pas fermer (ex: "Il reste de l'argent" ou "Dette en cours").

// 3. Le Test (Le Main)

//     Initialiser un compte vide.

//     Essayer de déposer 100 (ça devrait échouer car compte fermé).

//     Ouvrir le compte.

//     Déposer 150.

//     Emprunter 5000.

//     Rembourser la dette.

//     Fermer le compte.

//     Afficher l'état final du compte.
