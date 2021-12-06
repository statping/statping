/// <reference types="cypress" />

import "../support/commands"

context('Incidents Tests', () => {

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

  it('should create new incident', () => {
      cy.visit('/dashboard')
      cy.wait(3000)
      cy.get('.service_block').eq(0).find(".incident").click()
      cy.get('#title').clear().type('Downtime')
      cy.get('#description').clear().type('Recently we found an issue with authentication')
      cy.get('button[type="submit"]').click()
  })

})
