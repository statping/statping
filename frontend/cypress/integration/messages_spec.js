/// <reference types="cypress" />

import "../support/commands"

context('Announcements Tests', () => {


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

  it('should goto messages', () => {
    cy.visit('/dashboard/messages')
    cy.get('tbody > tr').should('have.length', 2)
  })

  it('should create Message', () => {
    cy.visit('/dashboard/messages')
    cy.get('#title').clear().type('Test Message')
    cy.get('#description').clear().type('This message was created by Cypress!')

    cy.get('button[type="submit"]').click()
  })

  it('should confirm new Message', () => {
    cy.visit('/dashboard/messages')
    cy.get('tbody > tr').should('have.length', 3)
  })

  it('should confirm delete Message', () => {
    cy.visit('/dashboard/messages')
    cy.get('tbody > tr').eq(0).find('.btn-danger').click()
    cy.get('tbody > tr').should('have.length', 2)
  })

})
