/// <reference types="cypress" />

import "../support/commands"

context('Notifier Tests', () => {


  beforeEach(() => {
    cy.restoreLocalStorageCache();
  });

  afterEach(() => {
    cy.saveLocalStorageCache();
  });

  it('should Login', () => {
    cy.visit('/login')
    cy.get('#username').clear().type('admin')
    cy.get('#password').clear().type('admin')
    cy.get('button[type="submit"]').click()

    cy.get('.navbar-brand').should('contain', 'Statping')
    cy.getCookies()

    cy.getCookies().should('have.length', 1)
  })

  it('should confirm notifiers are installed', () => {
    cy.visit('/dashboard/settings')
    cy.get('#notifiers_tabs > a').should('have.length', 9)

    cy.get('#api_key').should('not.have.value', '')
    cy.get('#api_secret').should('not.have.value', '')
  })

})
